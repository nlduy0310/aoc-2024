package maths

import "math"

func GCD(a, b int) int {

	for b != 0 {
		a, b = b, a%b
	}

	if a < 0 {
		return -a
	} else {
		return a
	}
}

func Reduce(a, b int) (int, int) {

	gcd := GCD(a, b)

	return a / gcd, b / gcd
}

func Digits(a int) int {

	if a == 0 {
		return 1
	}

	if a < 0 {
		a = -a
	}

	return int(math.Floor(math.Log10(float64(a)))) + 1
}
