package stone

import (
	"fmt"
	"math"

	maths_utils "github.com/nlduy0310/aoc-2024/utils/maths"
)

type Stone struct {
	Val int
}

func (s Stone) String() string {

	return fmt.Sprintf("Stone[Val=%d]", s.Val)
}

func (s Stone) Change() []Stone {

	if s.Val == 0 {
		return []Stone{{Val: 1}}
	}

	digitsCount := maths_utils.Digits(s.Val)
	if digitsCount%2 == 0 {
		firstHalf := s.Val / int(math.Pow10(digitsCount/2))
		secondHalf := s.Val % int(math.Pow10(digitsCount/2))
		return []Stone{{Val: firstHalf}, {Val: secondHalf}}
	}

	return []Stone{{Val: s.Val * 2024}}
}
