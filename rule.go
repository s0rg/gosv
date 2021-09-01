package gosv

import "strings"

type Rule func(*ruleSet)

func MinLen(val int) Rule {
	return func(r *ruleSet) {
		r.lenMin = val
	}
}

func MaxLen(val int) Rule {
	return func(r *ruleSet) {
		r.lenMax = val
	}
}

func MinUppers(val int) Rule {
	return func(r *ruleSet) {
		r.upsMin = val
	}
}

func MaxUppers(val int) Rule {
	return func(r *ruleSet) {
		r.upsMax = val
	}
}

func MinLowers(val int) Rule {
	return func(r *ruleSet) {
		r.lowMin = val
	}
}

func MaxLowers(val int) Rule {
	return func(r *ruleSet) {
		r.lowMax = val
	}
}

func MinNumbers(val int) Rule {
	return func(r *ruleSet) {
		r.numMin = val
	}
}

func MaxNumbers(val int) Rule {
	return func(r *ruleSet) {
		r.numMax = val
	}
}

func MaxSequencies(val float32) Rule {
	return func(r *ruleSet) {
		r.seqMax = val
	}
}

func MaxDuplicates(val float32) Rule {
	return func(r *ruleSet) {
		r.dupMax = val
	}
}

func ForbidContent(vals ...string) Rule {
	return func(r *ruleSet) {
		r.forbid = make([]string, len(vals))

		for i := 0; i < len(vals); i++ {
			r.forbid[i] = strings.ToLower(vals[i])
		}
	}
}
