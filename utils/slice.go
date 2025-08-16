package utils

func SliceCopy[T PrimitiveCopyable](source []T) []T {

	if source == nil {
		return nil
	}

	ret := make([]T, len(source))
	copy(ret, source)

	return ret
}

func SliceEqual[T comparable](first, second []T) bool {

	if first == nil && second == nil {
		return true
	} else if first == nil {
		return second == nil
	} else if second == nil {
		return first == nil
	} else if len(first) != len(second) {
		return false
	}

	for i := range len(first) {
		if first[i] != second[i] {
			return false
		}
	}

	return true
}

func SliceMap[T any, K any](source []T, mapperFunction func(T) K) []K {

	if source == nil {
		return nil
	}

	ret := make([]K, len(source))

	for idx, val := range source {
		ret[idx] = mapperFunction(val)
	}

	return ret
}
