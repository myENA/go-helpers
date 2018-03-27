package helpers

import (
	"strings"
	"sort"
)

// CleanStringSlice takes an input and will attempt to "clean" the input by removing "empty" or "whitespace-only"
// values.  You will always get a []string response, even if you pass in a nil
func CleanStringSlice(in []string) (out []string) {
	var inVal, outVal string

	// if in nil or empty...
	if nil == in || 0 == len(in) {
		return
	}

	for _, inVal = range in {
		outVal = strings.TrimSpace(inVal)
		if "" != outVal {
			out = append(out, outVal)
		}
	}

	return
}

// UniqueStringSlice takes a string array and returns a slice with duplicates removed
// This also strips "empty" values out.  You will always get a []string response, even if you pass in a nil
func UniqueStringSlice(in []string) (out []string) {
	var inVal, outVal string

	// if in nil or empty...
	if nil == in || 0 == len(in) {
		return
	}

UniqueLoop:
	for _, inVal = range in {
		// see if empty, but don't change values for non-empty
		inVal = strings.TrimSpace(inVal)
		if "" == inVal {
			continue UniqueLoop
		}

		for _, outVal = range out {
			if outVal == inVal {
				continue UniqueLoop
			}
		}

		out = append(out, inVal)
	}

	return
}

// CombineStringSlices takes 1..* string arrays and combines them into a single slice.
// This will also unique and strip "empty" values out, so y'know...there you go.
func CombineStringSlices(ins ...[]string) (out []string, delta int) {
	var i int
	var in []string
	var inVal, outVal string

	// if there was no input...
	if 0 == len(ins) {
		return
	}

	for i, in = range ins {
		// if "empty", just move on
		if nil == in || 0 == len(in) {
			continue
		}

	ValueLoop:
		for _, inVal = range UniqueStringSlice(in) {
			for _, outVal = range out {
				if outVal == inVal {
					continue ValueLoop
				}
			}

			out = append(out, inVal)
			if 0 < i {
				delta++
			}
		}
	}

	return
}

// RemoveStringsFromSlice will attempt to remove values present in "remove" from "root"
func RemoveStringsFromSlice(root, remove []string) (out []string, delta int) {
	var rootVal, removeVal string

	// if root is nil, just return.
	if nil == root {
		return
	}

	// if remove list empty, copy root to out
	if nil == remove || 0 == len(remove) {
		out = make([]string, len(root))
		copy(out, root)
		return
	}

RootLoop:
	for _, rootVal = range root {
		for _, removeVal = range remove {
			if rootVal == removeVal {
				delta++
				continue RootLoop
			}
		}
		out = append(out, rootVal)
	}

	return
}

// StringSlicesEqual will attempt to determine if both provided string slices contain the same values.  The original
// slices are not modified.
func StringSlicesEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	ac := make([]string, len(a))
	bc := make([]string, len(b))
	copy(ac, a)
	copy(bc, b)
	sort.Strings(ac)
	sort.Strings(bc)
	for i, v := range ac {
		if bc[i] != v {
			return false
		}
	}
	return true
}
