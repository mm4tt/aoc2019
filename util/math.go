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

func LMC(a ...int) int {
	lcm := a[0]
	for _, b := range a[1:] {
		lcm = lcm * b / GCD(lcm, b)
	}
	return lcm
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

func SgnI(a int) int {
	if a == 0 {
		return 0
	}
	return a / Abs(a)
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Ceil(a, b int) int {
	return int(math.Ceil(float64(a) / float64(b)))
}
