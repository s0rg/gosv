package gosv

type stats struct {
	Len          int
	Uppers       int
	Lowers       int
	Numbers      int
	Sequencies   int
	Duplicates   int
	Entropy      float64
	EntropyIdeal float64
}

func (s *stats) Equals(other *stats) (yes bool) {
	switch {
	case s.Len != other.Len:
		return
	case s.Uppers != other.Uppers:
		return
	case s.Lowers != other.Lowers:
		return
	case s.Numbers != other.Numbers:
		return
	case s.Sequencies != other.Sequencies:
		return
	case s.Duplicates != other.Duplicates:
		return
	}

	return true
}
