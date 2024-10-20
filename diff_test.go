package gotextdiff_test

import (
	_ "embed"
	"fmt"
	"strings"
	"testing"

	"github.com/OrlovPavel/gotextdiff"
	diff "github.com/OrlovPavel/gotextdiff"
	"github.com/OrlovPavel/gotextdiff/difftest"
	"github.com/OrlovPavel/gotextdiff/myers"
	"github.com/OrlovPavel/gotextdiff/span"
)

//go:embed test_resources/old.java.txt
var javaOld string

//go:embed test_resources/new.java.txt
var javaNew string

//go:embed test_resources/desired.txt
var javaDiff string

func TestApplyEdits(t *testing.T) {
	for _, tc := range difftest.TestCases {
		t.Run(tc.Name, func(t *testing.T) {
			t.Helper()
			if got := diff.ApplyEdits(tc.In, tc.Edits); got != tc.Out {
				t.Errorf("ApplyEdits edits got %q, want %q", got, tc.Out)
			}
			if tc.LineEdits != nil {
				if got := diff.ApplyEdits(tc.In, tc.LineEdits); got != tc.Out {
					t.Errorf("ApplyEdits lineEdits got %q, want %q", got, tc.Out)
				}
			}
		})
	}
}

func TestLineEdits(t *testing.T) {
	for _, tc := range difftest.TestCases {
		t.Run(tc.Name, func(t *testing.T) {
			t.Helper()
			// if line edits not specified, it is the same as edits
			edits := tc.LineEdits
			if edits == nil {
				edits = tc.Edits
			}
			if got := diff.LineEdits(tc.In, tc.Edits); diffEdits(got, edits) {
				t.Errorf("LineEdits got %q, want %q", got, edits)
			}
		})
	}
}

func TestUnified(t *testing.T) {
	for _, tc := range difftest.TestCases {
		t.Run(tc.Name, func(t *testing.T) {
			t.Helper()
			unified := fmt.Sprint(diff.ToUnified(difftest.FileA, difftest.FileB, tc.In, tc.Edits))
			if unified != tc.Unified {
				t.Errorf("edits got diff:\n%v\nexpected:\n%v", unified, tc.Unified)
			}
			if tc.LineEdits != nil {
				unified := fmt.Sprint(diff.ToUnified(difftest.FileA, difftest.FileB, tc.In, tc.LineEdits))
				if unified != tc.Unified {
					t.Errorf("lineEdits got diff:\n%v\nexpected:\n%v", unified, tc.Unified)
				}
			}
		})
	}
}

func TestUnifiedOutput(t *testing.T) {
	edits := myers.ComputeEdits(span.URIFromPath(""), javaOld, javaNew)
	diffRaw := fmt.Sprint(gotextdiff.ToUnified("", "", javaOld, edits))

	diff := strings.Join(strings.Split(diffRaw, "\n")[2:], "\n")

	if javaDiff != diff {
		t.Errorf("got diff:\n%v\nexpected:\n%v", diffRaw, javaDiff)
	}
}

func diffEdits(got, want []diff.TextEdit) bool {
	if len(got) != len(want) {
		return true
	}
	for i, w := range want {
		g := got[i]
		if span.Compare(w.Span, g.Span) != 0 {
			return true
		}
		if w.NewText != g.NewText {
			return true
		}
	}
	return false
}
