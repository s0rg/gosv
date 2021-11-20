package gosv

import "testing"

func TestEquals(t *testing.T) {
	t.Parallel()

	const val = 1

	var a, b stats

	a.Len = val

	if a.Equals(&b) {
		t.Error("step 1")
	}

	b.Len = val
	a.Uppers = val

	if a.Equals(&b) {
		t.Error("step 2")
	}

	b.Uppers = val
	a.Lowers = val

	if a.Equals(&b) {
		t.Error("step 3")
	}

	b.Lowers = val
	a.Numbers = val

	if a.Equals(&b) {
		t.Error("step 4")
	}

	b.Numbers = val
	a.Sequencies = val

	if a.Equals(&b) {
		t.Error("step 5")
	}

	b.Sequencies = val
	a.Duplicates = val

	if a.Equals(&b) {
		t.Error("step 6")
	}
}
