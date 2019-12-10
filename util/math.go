package util

import "math"

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return Abs(a)
}

func Abs(a int) int {
	if a < 0 {
		a = -a
	}
	return a
}

func Sgn(a float64) float64 {
	if a == 0 {
		return 0
	}
	return a / math.Abs(a)
}
