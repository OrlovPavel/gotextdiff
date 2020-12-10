# gotextdiff

This is a copy of the Go text diffing package that [the official Go language server gopls uses internally](https://github.com/golang/tools/tree/master/internal/lsp/diff).

This is arguably one of the best (and most maintained) text diffing packages in Go as of at least 2020.

## Alternatives

- [github.com/sergi/go-diff](https://github.com/sergi/go-diff): A Go language port of Neil Fraser's google-diff-match-patch code
- [github.com/andreyvit/diff](https://github.com/andreyvit/diff): Quick'n'easy string diffing functions for Golang based on github.com/sergi/go-diff.
- [github.com/kylelemons/godebug/diff](https://github.com/kylelemons/godebug/tree/master/diff): implements a linewise diff algorithm ([inactive](https://github.com/kylelemons/godebug/issues/22#issuecomment-524573477)).

## Example usage

Assuming you want to diff `a.txt` and `b.txt`, whose contents are stored in `aString` and `bString` then:

```Go
edits := myers.ComputeEdits(span.URIFromPath("a.txt"), aString, bString)
diff := fmt.Sprint(gotextdiff.ToUnified("a.txt", "b.txt", aString, edits))
```

`diff` will be a string like:

```diff
--- a.txt
+++ b.txt
@@ -1,13 +1,28 @@
-foo
+bar
```

## API compatability

We will publish a new major version anytime the API changes in a backwards-incompatible way. Because the upstream is not being developed with this being a public package in mind, API breakages may occur more often than in other Go packages.

## Contributing

We will only accept changes made upstream, please send any contributions to the upstream instead!

## License

See https://github.com/golang/tools/blob/master/LICENSE

