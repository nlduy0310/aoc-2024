package main

type CopySafePrimitive interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~complex64 | ~complex128 |
		~bool | ~string |
		RulePair // :D
}

func panicIf(err error) {

	if err != nil {
		panic(err)
	}
}

func safeCopyList[T CopySafePrimitive](src []T) []T {

	ret := make([]T, len(src))
	copy(ret, src)

	return ret
}

func safeCopyMap[K CopySafePrimitive, V CopySafePrimitive](src map[K]V) map[K]V {

	ret := make(map[K]V, len(src))

	for key, value := range src {
		ret[key] = value
	}

	return ret
}

func contains[T comparable](slice []T, element T) bool {

	for _, val := range slice {
		if val == element {
			return true
		}
	}

	return false
}

func count[T comparable](slice []T, element T) int {

	ret := 0

	for _, val := range slice {
		if val == element {
			ret += 1
		}
	}

	return ret
}

func mapTo[T any, K any](source []T, transformer func(T) K) []K {

	if source == nil {
		panic("can not map nil value")
	}

	ret := make([]K, len(source))

	for idx, val := range source {
		ret[idx] = transformer(val)
	}

	return ret
}

func filter[T CopySafePrimitive](source []T, filter func(T) bool) []T {

	if source == nil {
		panic("can not filter nil value")
	}

	ret := make([]T, 0)

	for _, val := range source {
		if filter(val) {
			ret = append(ret, val)
		}
	}

	return ret
}

func reserve[T CopySafePrimitive](source []T, reserveSlice []T) []T {

	ret := make([]T, 0, len(reserveSlice))

	for _, val := range source {
		if contains(reserveSlice, val) {
			ret = append(ret, val)
		}
	}

	return ret
}

// func logVar[T any](name string, variable T) {
// 	fmt.Printf("%s: %v\n", name, variable)
// }
