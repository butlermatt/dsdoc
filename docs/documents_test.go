package docs

import "testing"

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
