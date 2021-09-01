package gosv

import "unicode"

func calcStats(input string) (s stats) {
	var (
		prev               = rune(-1)
		isChar, prevIsChar bool
	)

	for _, c := range input {
		s.Len++

		switch {
		case unicode.IsUpper(c):
			s.Uppers++

			c = unicode.ToLower(c)
			isChar = true
		case unicode.IsLower(c):
			s.Lowers++

			isChar = true
		case unicode.IsNumber(c):
			s.Numbers++

			isChar = true
		}

		switch {
		case prev < 0:
		case prev == c:
			s.Duplicates++
		case (isChar && prevIsChar) && isNeighbours(prev, c):
			s.Sequencies++
		}

		prev, prevIsChar, isChar = c, isChar, false
	}

	return s
}

func isNeighbours(a, b rune) bool {
	return abs32(a-b) == 1
}

func abs32(i int32) int32 {
	if i < 0 {
		i = -i
	}

	return i
}
