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

func TrimDsDoc(s []string) ([]string, int) {
	var r []string

	var found bool
	var i int
	var str string
	for i, str = range s {
		if j := strings.Index(str, Prefix); j != -1 {
			found = true
			r = append(r, str[j:])
		} else {
			if found {
				break
			}
		}
	}

	return r, i
}
