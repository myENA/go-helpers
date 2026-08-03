# go-helpers

A small collection of Go helper utilities for working with string slices.

[![Go Reference](https://pkg.go.dev/badge/github.com/myENA/go-helpers.svg)](https://pkg.go.dev/github.com/myENA/go-helpers)

## Installation

```bash
go get github.com/myENA/go-helpers
```

## Usage

```go
import "github.com/myENA/go-helpers"
```

### CleanStringSlice

Trims whitespace and removes empty or whitespace-only values. Always returns a non-nil `[]string`, even when the input is `nil`.

```go
helpers.CleanStringSlice([]string{"ok", " extralead", "extrafollow ", "       "})
// []string{"ok", "extralead", "extrafollow"}
```

### UniqueStringSlice

Removes duplicates and empty/whitespace-only values. Trimmed values are compared for uniqueness; non-empty results keep their trimmed form.

```go
helpers.UniqueStringSlice([]string{"ok", "dupe1", "dupe2", "     dupe2", "dupe1 ", "ok2"})
// []string{"ok", "dupe1", "dupe2", "ok2"}
```

### CombineStringSlices

Merges one or more string slices into a single unique slice (empty values excluded). Returns the combined slice and a `delta` count of values added from slices after the first.

```go
combined, delta := helpers.CombineStringSlices(
    []string{"ok", "ok", "val1", "val2", "      val3"},
    []string{"ok2", "ok", "val3", "val2"},
)
// combined: []string{"ok", "val1", "val2", "val3", "ok2"}
// delta: 1
```

### RemoveStringsFromSlice

Removes values present in `remove` from `root`. Returns the resulting slice and the number of removed entries. When `remove` is empty, returns a copy of `root`.

```go
out, delta := helpers.RemoveStringsFromSlice(
    []string{"ok", "removeme1", "removeme2", "ok2"},
    []string{"removeme1", "removeme2"},
)
// out: []string{"ok", "ok2"}
// delta: 2
```

### StringSlicesEqual

Reports whether two string slices contain the same values, ignoring order. The original slices are not modified.

```go
helpers.StringSlicesEqual([]string{"b", "a"}, []string{"a", "b"}) // true
```

## License

Mozilla Public License Version 2.0. See [LICENSE](LICENSE).
