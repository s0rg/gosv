package gosv

// Validate runs set of checks against input string,
// returns nil if all checks are passed.
func Validate(input string, rules ...Rule) (err error) {
	var rls ruleSet

	return rls.Add(rules...).Validate(input)
}
