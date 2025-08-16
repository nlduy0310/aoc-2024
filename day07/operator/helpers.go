package operator

import "math"

func digits(n int) int {

	if n == 0 {
		return 1
	}

	if n < 0 {
		n = -n
	}

	return int(math.Floor(math.Log10(float64(n)))) + 1
}
