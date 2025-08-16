package maths

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
