package trim

import (
	"strings"
)

// Prefix is the DsDoc comment style.
const Prefix string = "//*"

func TrimDsDoc(s []string) [][]string {
	var r [][]string
	var b []string

	var found bool
	var str string
	for _, str = range s {
		if j := strings.Index(str, Prefix); j != -1 {
			found = true
			b = append(b, strings.TrimSpace(str[(j + len(Prefix)):]))
		} else {
			if found {
				tmp := make([]string, len(b))
				copy(tmp, b)
				r = append(r, tmp)
				b = nil
				found = false
			}
		}
	}
	if len(b) > 0 {
		tmp := make([]string, len(b))
		copy(tmp, b)
		r = append(r, tmp)
	}

	return r
}
