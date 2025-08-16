package operator

import (
	"fmt"
	"math"

	"github.com/nlduy0310/aoc-2024/utils"
)

type Operator int

const (
	OpAdd Operator = iota
	OpSubtract
	OpMultiply
	OpDivide
	OpConcatenate
)

var validOperators = map[Operator]struct{}{
	OpAdd:         {},
	OpSubtract:    {},
	OpMultiply:    {},
	OpDivide:      {},
	OpConcatenate: {},
}

var applyFuncMap = map[Operator]func(int, int) int{
	OpAdd:         func(a, b int) int { return a + b },
	OpSubtract:    func(a, b int) int { return a - b },
	OpMultiply:    func(a, b int) int { return a * b },
	OpDivide:      func(a, b int) int { return a / b },
	OpConcatenate: func(a, b int) int { return a*int(math.Pow10(digits(b))) + b },
}

func ensureValid(o Operator) {

	utils.Assert(o.IsValid(), fmt.Sprintf("invalid operator enum %d", o))
}

func (o Operator) IsValid() bool {

	_, valid := validOperators[o]

	return valid
}

func (o Operator) String() string {

	ensureValid(o)
	var opString string = "unknown"

	switch o {
	case OpAdd:
		opString = "+"
	case OpSubtract:
		opString = "-"
	case OpMultiply:
		opString = "*"
	case OpDivide:
		opString = "/"
	case OpConcatenate:
		opString = "||"
	}

	return fmt.Sprintf("Operator[%s]", opString)
}

// For some reason, this can not be defined with generics
func (o Operator) Apply(a, b int) int {

	ensureValid(o)

	applyFunc := applyFuncMap[o]

	return applyFunc(a, b)
}
