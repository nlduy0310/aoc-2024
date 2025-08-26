package position

import (
	"fmt"
	"math"
)

func clampValue(val, rangeMin, rangeMax int) (int, error) {

	if rangeMin > rangeMax {
		return 0, fmt.Errorf("min value (%d) > max value (%d)", rangeMin, rangeMax)
	}

	if val >= rangeMin && val <= rangeMax {
		return val, nil
	}

	var valRange int = rangeMax - rangeMin + 1
	var diff, offset int

	diff = min(int(math.Abs(float64(val-rangeMin))), int(math.Abs(float64(val-rangeMax))))
	offset = diff % valRange

	if offset == 0 {
		offset = valRange - 1
	} else {
		offset = offset - 1
	}

	if val < rangeMin {
		return rangeMax - offset, nil
	} else {
		return rangeMin + offset, nil
	}
}
