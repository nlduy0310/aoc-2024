package utils

import (
	"fmt"
	"slices"
)

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

// Returns the index of the first element to cause evalFunc to returns true. If none are found, returns -1.
func SliceFindFirst[T any](slice []T, evalFunc func(T) bool) int {

	for idx, t := range slice {
		if evalFunc(t) {
			return idx
		}
	}

	return -1
}

// Returns the index of the last element to cause evalFunc to returns true. If none are found, returns -1.
func SliceFindLast[T any](slice []T, evalFunc func(T) bool) int {

	for idx := len(slice) - 1; idx >= 0; idx-- {
		if evalFunc(slice[idx]) {
			return idx
		}
	}

	return -1
}

func SliceInit[T any](length int, defaultValue T) []T {

	Assert(length >= 0, fmt.Sprintf("invalid length %d", length))

	ret := make([]T, length)

	for idx := range length {
		ret[idx] = defaultValue
	}

	return ret
}

// Returns a copy of the filtered slice
func SliceFilter[T any](source []T, keep func(T) bool) []T {

	ret := make([]T, 0, len(source))

	for i := range len(source) {
		if keep(source[i]) {
			ret = append(ret, source[i])
		}
	}

	return slices.Clip(ret)
}
