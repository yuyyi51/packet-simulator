package utils

import "math/rand"

type Fraction struct {
	numerator   int64
	denominator int64
}

func NewFraction(numerator, denominator int64) Fraction {
	return Fraction{
		numerator:   numerator,
		denominator: denominator,
	}
}

func (f Fraction) Calculate() float64 {
	return float64(f.numerator) / float64(f.denominator)
}

func (f Fraction) Rand(r *rand.Rand) bool {
	if f.denominator <= 0 {
		return false
	}
	return r.Int63n(f.denominator) < f.numerator
}
