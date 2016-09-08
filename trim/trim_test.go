package trim

import (
	"testing"
)

// Ensure that comments can be at the end of text/source.
func TestTrimDsDoc(t *testing.T) {
	var tests = []struct {
		in  []string
		out [][]string
	}{
		{
			in: []string{
				`some text //* line one`,
				`some text //* line two`,
				`some text //* `,
				`some text //* line four`,
			},
			out: [][]string{
				{
					`line one`,
					`line two`,
					``,
					`line four`,
				},
			},
		},
		{
			in: []string{
				`some text //* line one`,
				`some text //* line two`,
				`some text`,
				`some text //* line four`,
				`some text //*`,
				`some text //* line five`,
			},
			out: [][]string{
				{
					`line one`,
					`line two`,
				},
				{
					`line four`,
					``,
					`line five`,
				},
			},
		},
		{
			in: []string{
				"some text",
				"some text ",
				"some text\t",
				"some text //* line one",
				"some text //* line two",
			},
			out: [][]string{
				{
					"line one",
					"line two",
				},
			},
		},
		{
			in: []string{
				"some text",
				"some text",
				"some text",
				"some text",
			},
			out: [][]string{},
		},
	}

	for i, tt := range tests {
		res := TrimDsDoc(tt.in)

		if len(res) != len(tt.out) {
			t.Fatalf("%d. Batch counts do not match: exp=%d got=%d", i, len(tt.out), len(res))
		}

		for j, bt := range tt.out {
			if len(bt) != len(res[j]) {
				t.Errorf("%d. %d. Line counts do not match: exp=%d got=%d", i, j, len(bt), len(res[j]))
			}
			for k, str := range bt {
				if str != res[j][k] {
					t.Errorf("%d. %q does not match %q", i, str, res[j][k])
				}
			}
		}
	}
}
