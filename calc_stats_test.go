package gosv

import "testing"

func TestCalc(t *testing.T) {
	t.Parallel()

	type testCase struct {
		Value string
		Want  *stats
	}

	cases := []testCase{
		{Value: "a", Want: &stats{Len: 1, Lowers: 1}},
		{Value: "B", Want: &stats{Len: 1, Uppers: 1}},
		{Value: "0", Want: &stats{Len: 1, Numbers: 1}},
		{Value: "aB", Want: &stats{Len: 2, Lowers: 1, Uppers: 1, Sequencies: 1}},
		{Value: "bA", Want: &stats{Len: 2, Lowers: 1, Uppers: 1, Sequencies: 1}},
		{Value: "bAa", Want: &stats{Len: 3, Lowers: 2, Uppers: 1, Sequencies: 1, Duplicates: 1}},
		{Value: "aCd134", Want: &stats{Len: 6, Lowers: 2, Uppers: 1, Numbers: 3, Sequencies: 2}},
		{Value: "aaaabc", Want: &stats{Len: 6, Lowers: 6, Sequencies: 2, Duplicates: 3}},
		{Value: "aaaaec", Want: &stats{Len: 6, Lowers: 6, Duplicates: 3}},
		{Value: "aD3x51X[.!", Want: &stats{Len: 10, Lowers: 2, Uppers: 2, Numbers: 3}},
	}

	for i, c := range cases {
		s := calcStats(c.Value)

		if !s.Equals(c.Want) {
			t.Fatalf("case %d failed for: '%s' got: %+v want: %+v", i, c.Value, s, c.Want)
		}
	}
}
