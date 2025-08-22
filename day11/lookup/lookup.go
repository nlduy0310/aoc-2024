package lookup

type Lookup struct {
	cache map[int]map[int]int
}

func NewLookup() Lookup {

	return Lookup{cache: make(map[int]map[int]int)}
}

func (l Lookup) Mark(stoneValue, depth, count int) {

	if _, ok := l.cache[stoneValue]; !ok {
		l.cache[stoneValue] = make(map[int]int)
	}

	l.cache[stoneValue][depth] = count
}

func (l Lookup) TryGet(stoneValue, depth int) (int, bool) {

	current, ok := l.cache[stoneValue]
	if !ok {
		return 0, false
	}

	ret, found := current[depth]
	return ret, found
}
