package trim

import (
	"testing"
)

// Ensure the DSDoc comment is stripped
func TestTrimPrefix(t *testing.T) {
	sample := []string{
		`//* line one`,
		`//* line two`,
	}

	expect := []string{
		`line one`,
		`line two`,
	}

	res, err := TrimPrefix(sample)
	if err != nil {
		t.Fatal("Unexpected error occurred", err)
	}
	for i, str := range res {
		if str != expect[i] {
			t.Fatalf("%s does not match %s", str, expect[i])
		}
	}
}

// Expect an error return if non-dsdoc comment line is included
func TestTrimPrefix2(t *testing.T) {
	sample := []string{
		`//* line one`,
		`no comment`,
		`//* line two`,
	}

	_, err := TrimPrefix(sample)
	if err == nil {
		t.Fatal("No error when an error was expected")
	}
}

// Ensure empty lines are preserved.
func TestTrimPrefix3(t *testing.T) {
	sample := []string{
		`//* line one`,
		`//*         `,
		`//* line three`,
		`//*`,
	}

	expect := []string{
		`line one`,
		``,
		`line three`,
		``,
	}

	res, err := TrimPrefix(sample)
	if err != nil {
		t.Fatalf("Unexpected error occurred: %v", err)
	}
	for i, str := range res {
		if str != expect[i] {
			t.Fatalf("%s does not match %s", str, expect[i])
		}
	}
}

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
					`//* line one`,
					`//* line two`,
					`//* `,
					`//* line four`,
				},
			},
		},
		{
			in: []string{
				`some text //* line one`,
				`some text //* line two`,
				`some text`,
				`some text //* line four`,
			},
			out: [][]string{
				{
					`//* line one`,
					`//* line two`,
				},
				{
					`//* line four`,
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
					"//* line one",
					"//* line two",
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
