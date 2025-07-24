package utils

func IsInRangeInclusive[T Ordered](value, rangeStart, rangeEnd T) bool {

	return value >= rangeStart && value <= rangeEnd
}

func IsInRangeExclusive[T Ordered](value, rangeStart, rangeEnd T) bool {

	return value > rangeStart && value < rangeEnd
}

func MinMax[T Ordered](firstValue, secondValue T) (T, T) {

	if firstValue < secondValue {
		return firstValue, secondValue
	}

	return secondValue, firstValue
}
