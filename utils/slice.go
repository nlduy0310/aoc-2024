package utils

func SliceCopy[T PrimitiveCopyable](source []T) []T {

	if source == nil {
		return nil
	}

	ret := make([]T, len(source))
	copy(ret, source)

	return ret
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
