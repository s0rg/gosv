package gosv

import "strings"

// Rule for validation.
type Rule func(*ruleSet)

// MinLen returns rule with given min string length.
func MinLen(val int) Rule {
	return func(r *ruleSet) {
		r.lenMin = val
	}
}

// MaxLen returns rule with given max string length.
func MaxLen(val int) Rule {
	return func(r *ruleSet) {
		r.lenMax = val
	}
}

// MinUppers returns rule with min uppercase chars count.
func MinUppers(val int) Rule {
	return func(r *ruleSet) {
		r.upsMin = val
	}
}

// MaxUppers returns rule with max uppercase chars count.
func MaxUppers(val int) Rule {
	return func(r *ruleSet) {
		r.upsMax = val
	}
}

// MinLowers returns rule with min lowercase chars count.
func MinLowers(val int) Rule {
	return func(r *ruleSet) {
		r.lowMin = val
	}
}

// MaxLowers returns rule with max lowercase chars count.
func MaxLowers(val int) Rule {
	return func(r *ruleSet) {
		r.lowMax = val
	}
}

// MinNumbers returns rule with min numbers count.
func MinNumbers(val int) Rule {
	return func(r *ruleSet) {
		r.numMin = val
	}
}

// MaxNumbers returns rule with max numbers count.
func MaxNumbers(val int) Rule {
	return func(r *ruleSet) {
		r.numMax = val
	}
}

// MaxSequencies returns rule with max sequencies (i.e. '123', 'abc') rate.
func MaxSequencies(val float32) Rule {
	return func(r *ruleSet) {
		r.seqMax = val
	}
}

// MaxDuplicates returns rule with max dulpicates rate.
func MaxDuplicates(val float32) Rule {
	return func(r *ruleSet) {
		r.dupMax = val
	}
}

// ForbidContent returns rule with set of forbidden strings.
func ForbidContent(vals ...string) Rule {
	return func(r *ruleSet) {
		r.forbid = make([]string, len(vals))

		for i := 0; i < len(vals); i++ {
			r.forbid[i] = strings.ToLower(vals[i])
		}
	}
}

// MaxEntropyDiff returns rule with max entropy diff between ideal and current values.
func MaxEntropyDiff(val float64) Rule {
	return func(r *ruleSet) {
		r.maxEntropyDiff = val
	}
}
