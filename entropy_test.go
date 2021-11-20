package gosv

import (
	"math"
	"testing"
)

type runeAdder interface {
	Add(rune)
}

func feedString(val string, ra runeAdder) {
	for _, r := range val {
		ra.Add(r)
	}
}

func equalFloats(a, b float64) bool {
	const epsilon = float64(0.00001)

	return math.Abs(a-b) < epsilon
}

func TestShannonEntropy(t *testing.T) {
	t.Parallel()

	type testCase struct {
		Val  string
		Ent1 float64
		Ent2 float64
	}

	cases := []testCase{
		{"abcd", 1.38629, 1.38629},
		{"aaaa", 0.00000, 1.38629},
		{"qwer", 1.38629, 1.38629},
		{"q1w2", 1.38629, 1.38629},
		{"1234", 1.38629, 1.38629},
		{"Hello world!", 2.09473, 2.48491},
	}

	for i, tc := range cases {
		e := shannonEntropy()

		feedString(tc.Val, e)

		r1 := e.Sum()
		r2 := e.SumIdeal()

		if !equalFloats(r1, tc.Ent1) {
			t.Errorf("case: %d wrong sum: want: %.05f got: %.05f", i, tc.Ent1, r1)
		}

		if !equalFloats(e.SumIdeal(), tc.Ent2) {
			t.Errorf("case: %d wrong ideal: want: %.05f got: %.05f", i, tc.Ent2, r2)
		}
	}
}
