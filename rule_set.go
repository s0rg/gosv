package gosv

import (
	"errors"
	"strings"
)

var (
	// ErrBadLength returned when length is out of a given range.
	ErrBadLength = errors.New("bad length")
	// ErrBadCounts returned when uppers/lowers/numbers rates out of a given range.
	ErrBadCounts = errors.New("bad counts")
	// ErrBadSeqRate returned when too many sequencies (ie '123' or 'abc').
	ErrBadSeqRate = errors.New("bad seq rate")
	// ErrBadDupRate returned when too many duplicates.
	ErrBadDupRate = errors.New("bad dup rate")
	// ErrHasForbidden returned when some forbidden words found.
	ErrHasForbidden = errors.New("has forbidden")
)

type ruleSet struct {
	lenMin, lenMax int
	upsMin, upsMax int
	lowMin, lowMax int
	numMin, numMax int
	seqMax, dupMax float32
	forbid         []string
}

func calcRate(a, b int) (c float32) {
	return float32(a) / float32(b)
}

func (r *ruleSet) Validate(input string) (err error) {
	s := calcStats(input)

	if !r.validLen(s.Len) {
		return ErrBadLength
	}

	if err = r.validRates(&s); err != nil {
		return err
	}

	if !r.checkForbidden(input) {
		return ErrHasForbidden
	}

	return nil
}

func (r *ruleSet) validRates(s *stats) (err error) {
	if !r.validLowers(s.Lowers) || !r.validUppers(s.Uppers) || !r.validNumbers(s.Numbers) {
		return ErrBadCounts
	}

	if !r.validSeqRate(calcRate(s.Sequencies, s.Len)) {
		return ErrBadSeqRate
	}

	if !r.validDupRate(calcRate(s.Duplicates, s.Len)) {
		return ErrBadDupRate
	}

	return nil
}

// Add rule to set.
func (r *ruleSet) Add(rules ...Rule) *ruleSet {
	for _, rfn := range rules {
		rfn(r)
	}

	return r
}

func (r *ruleSet) validLen(val int) bool {
	if r.lenMin > 0 && val < r.lenMin {
		return false
	}

	if r.lenMax > 0 && val > r.lenMax {
		return false
	}

	return true
}

func (r *ruleSet) validUppers(val int) bool {
	if r.upsMin > 0 && val < r.upsMin {
		return false
	}

	if r.upsMax > 0 && val > r.upsMax {
		return false
	}

	return true
}

func (r *ruleSet) validLowers(val int) bool {
	if r.lowMin > 0 && val < r.lowMin {
		return false
	}

	if r.lowMax > 0 && val > r.lowMax {
		return false
	}

	return true
}

func (r *ruleSet) validNumbers(val int) bool {
	if r.numMin > 0 && val < r.numMin {
		return false
	}

	if r.numMax > 0 && val > r.numMax {
		return false
	}

	return true
}

func (r *ruleSet) validSeqRate(val float32) bool {
	return val <= r.seqMax
}

func (r *ruleSet) validDupRate(val float32) bool {
	return val <= r.dupMax
}

func (r *ruleSet) checkForbidden(val string) bool {
	lval := strings.ToLower(val)

	for i := 0; i < len(r.forbid); i++ {
		if strings.Contains(lval, r.forbid[i]) {
			return false
		}
	}

	return true
}
