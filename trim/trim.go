package trim

import (
	"errors"
	"strings"
)

// Prefix is the DsDoc comment style.
const Prefix string = "//*"

// TrimPrefix takes the slice of strings which contain the DsDoc comment.
// It will skip any lines which do not have
func TrimPrefix(s []string) ([]string, error) {
	r := []string{}
	for _, str := range s {
		if !strings.HasPrefix(str, Prefix) {
			return nil, errors.New("Invalid DsDoc. Cannot contain a line without a DsDoc comment")
		}
		tmp := strings.TrimSpace(strings.Trim(str, Prefix))
		r = append(r, tmp)
	}

	return r, nil
}

func TrimDsDoc(s []string) [][]string {
	var r [][]string
	var b []string

	var found bool
	var str string
	for _, str = range s {
		if j := strings.Index(str, Prefix); j != -1 {
			found = true
			b = append(b, str[j:])
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
