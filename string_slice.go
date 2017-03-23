package helpers

import "strings"

// CleanStringSlice takes an input and will attempt to "clean" the input by removing "empty" or "whitespace-only"
// values.  You will always get a []string response, even if you pass in a nil
func CleanStringSlice(in []string) (out []string) {
	var inVal, outVal string

	if nil == in {
		out = make([]string, 0)
	} else if 0 == len(in) {
		out = in
	} else {
		for _, inVal = range in {
			outVal = strings.TrimSpace(inVal)
			if "" != outVal {
				out = append(out, outVal)
			}
		}
	}

	return
}

// UniqueStringSlice takes a string array and returns a slice with duplicates removed
// This also strips "empty" values out.  You will always get a []string response, even if you pass in a nil
func UniqueStringSlice(in []string) (out []string) {
	var inVal, outVal string

	if nil == in {
		out = make([]string, 0)
	} else if 0 == len(in) {
		out = in
	} else {
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
	}

	return
}

// CombineStringSlices takes 1..* string arrays and combines them into a single slice.
// This will also unique and strip "empty" values out, so y'know...there you go.
func CombineStringSlices(ins ...[]string) (out []string, additions int) {
	var i int
	var in []string
	var inVal, outVal string

	if 0 == len(ins) {
		out = make([]string, 0)
	} else {
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
					additions++
				}
			}
		}
	}

	return
}

// RemoveStringsFromSlice will attempt to remove values present in "remove" from "root"
func RemoveStringsFromSlice(root, remove []string) (out []string, removed int) {
	var rootVal, removeVal string

	if nil == root {
		out = make([]string, 0)
	} else if 0 == len(remove) {
		out = root
	} else {
	RootLoop:
		for _, rootVal = range root {
			for _, removeVal = range remove {
				if rootVal == removeVal {
					removed++
					continue RootLoop
				}
			}
			out = append(out, rootVal)
		}
	}

	return
}
