package gosv

import "math"

type entropy struct {
	chars  map[rune]int
	length int
}

func shannonEntropy() *entropy {
	return &entropy{
		length: 0,
		chars:  make(map[rune]int),
	}
}

func (e *entropy) Add(r rune) {
	e.chars[r]++
	e.length++
}

func (e *entropy) Sum() (rv float64) {
	for _, count := range e.chars {
		t := float64(count) / float64(e.length)
		rv += t * math.Log(t)
	}

	return -rv
}

func (e *entropy) SumIdeal() (rv float64) {
	p := float64(1.0) / float64(e.length)

	return -math.Log(p)
}
