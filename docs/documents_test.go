package docs

import "testing"

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
	sample := []string {
		`//* line one`,
		`//*         `,
		`//* line three`,
		`//*`,
	}

	expect := []string {
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
