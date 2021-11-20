package gosv

import (
	"testing"
)

func TestValidatePass(t *testing.T) {
	t.Parallel()

	cases := []string{
		"fo0B4r79",
		"dkxbkXq^97fZK",
		"p-+fSQz3B7g2g8-X",
		"2Q+d*4hJCe4Wnx9D",
		"57*g-V@wCq8d^_A",
	}

	rules := []Rule{
		MinLen(8),
		MaxLen(64),
		MinLowers(1),
		MinUppers(1),
		MinNumbers(1),
	}

	for i, c := range cases {
		if Validate(c, rules...) != nil {
			t.Fatalf("case: %d fail for: '%s'", i, c)
		}
	}
}

func TestValidateFail(t *testing.T) {
	t.Parallel()

	cases := []string{
		"aC3",
		"acE3",
		"acE3",
		"acEF3",
		"acEF35",      // ok but less than 8
		"abEF35Dg7$!", // ok but max len
	}

	rules := []Rule{
		MinLen(8),
		MaxLen(10),
		MinLowers(2),
		MinUppers(2),
		MinNumbers(2),
	}

	for i, c := range cases {
		if Validate(c, rules...) == nil {
			t.Fatalf("case: %d fail for: '%s'", i, c)
		}
	}
}

func TestValidateUppers(t *testing.T) {
	t.Parallel()

	cases := []string{
		"adg",
		"357",
		"25A",
		"4dN",
		"AGYN",
	}

	rules := []Rule{
		MinLen(3),
		MaxLen(4),
		MinUppers(2),
		MaxUppers(3),
	}

	for i, c := range cases {
		if Validate(c, rules...) == nil {
			t.Fatalf("case: %d fail for: '%s'", i, c)
		}
	}
}

func TestValidateLowers(t *testing.T) {
	t.Parallel()

	cases := []string{
		"ADG",
		"357",
		"25a",
		"4dN",
		"adhk",
	}

	rules := []Rule{
		MinLen(3),
		MaxLen(4),
		MinLowers(2),
		MaxLowers(3),
	}

	for i, c := range cases {
		if Validate(c, rules...) == nil {
			t.Fatalf("case: %d fail for: '%s'", i, c)
		}
	}
}

func TestValidateNumbers(t *testing.T) {
	t.Parallel()

	cases := []string{
		"ADG",
		"3cA",
		"2Da",
		"adgh",
		"1364",
	}

	rules := []Rule{
		MinLen(3),
		MaxLen(4),
		MinNumbers(2),
		MaxNumbers(3),
	}

	for i, c := range cases {
		if Validate(c, rules...) == nil {
			t.Fatalf("case: %d fail for: '%s'", i, c)
		}
	}
}

func TestValidateDupRate(t *testing.T) {
	t.Parallel()

	cases := []string{
		"aaQW2f",
		"aaaW2F",
	}

	rules := []Rule{
		MinLen(4),
		MaxLen(6),
		MaxDuplicates(0.4),
	}

	for i, c := range cases {
		if Validate(c, rules...) != nil {
			t.Fatalf("case: %d fail for: '%s'", i, c)
		}
	}

	if Validate("aaaaZ4", rules...) == nil {
		t.Fatal("last case fail")
	}
}

func TestValidateSequencies(t *testing.T) {
	t.Parallel()

	cases := []string{
		"abcd",
		"1234",
		"6789",
		"xyz1",
	}

	rules := []Rule{
		MinLen(4),
		MaxLen(10),
		MaxSequencies(0.4),
	}

	for i, c := range cases {
		if Validate(c, rules...) == nil {
			t.Fatalf("case: %d fail for: '%s'", i, c)
		}
	}

	const valid = "123sgo" // exactly 0.5

	if Validate(valid, rules...) != nil {
		t.Fatal("last case fail")
	}
}

func TestValidateForbidden(t *testing.T) {
	t.Parallel()

	cases := []string{
		"13forbidA",
		"forbid158",
	}

	rules := []Rule{
		MinLen(4),
		MaxLen(10),
		ForbidContent("forbid"),
	}

	for i, c := range cases {
		if Validate(c, rules...) == nil {
			t.Fatalf("case: %d fail for: '%s'", i, c)
		}
	}
}

func TestValidateEntropy(t *testing.T) {
	t.Parallel()

	cases := []string{
		"13f",
		"Hello world!",
	}

	rules := []Rule{
		MinLen(4),
		MaxEntropyDiff(0.3),
	}

	for i, c := range cases {
		if err := Validate(c, rules...); err == nil {
			t.Fatalf("case: %d fail for: '%s'", i, c)
		}
	}
}
