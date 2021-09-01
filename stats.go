package gosv

type stats struct {
	Len        int
	Uppers     int
	Lowers     int
	Numbers    int
	Sequencies int
	Duplicates int
}

func (s *stats) Equals(other *stats) bool {
	return s.Len == other.Len &&
		s.Uppers == other.Uppers &&
		s.Lowers == other.Lowers &&
		s.Numbers == other.Numbers &&
		s.Sequencies == other.Sequencies &&
		s.Duplicates == other.Duplicates
}
