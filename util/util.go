package util

func Abs(a int) uint {
	if a < 0 {
		a = -a
	}
	return uint(a)
}
