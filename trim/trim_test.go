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
	var tests = []struct{
		in  []string
		out []string
		ind int
	}{
		{
			in: []string{
				`some text //* line one`,
				`some text //* line two`,
				`some text //* `,
				`some text //* line four`,
			},
			out: []string{
				`//* line one`,
				`//* line two`,
				`//* `,
				`//* line four`,
			},
			ind: 3,
		},
		{
			in: []string{
				`some text //* line one`,
				`some text //* line two`,
				`some text`,
				`some text //* line four`,
			},
			out: []string{
				`//* line one`,
				`//* line two`,
			},
			ind: 2,
		},
		{
			in: []string{
				"some text",
				"some text ",
				"some text\t",
				"some text //* line one",
				"some text //* line two",
			},
			out: []string{
				"//* line one",
				"//* line two",
			},
			ind: 4,
		},
	}

	for i, tt := range tests {
		res, ind := TrimDsDoc(tt.in)

		if tt.ind != ind {
			t.Errorf("%d. Return index does not match: exp=%d got=%d", i, tt.ind, ind)
		}

		if len(res) != len(tt.out) {
			t.Errorf("%d. Line counts do not match: exp=%d got=%d", i, len(tt.out), len(res))
		}

		for j, str := range tt.out {
			if str != res[j] {
				t.Errorf("%d. %q does not match %q", i, str, res[j])
			}
		}
	}
}
