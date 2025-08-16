package equation

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nlduy0310/aoc-2024/utils"
)

type Equation struct {
	TestValue int
	Numbers   []int
}

func (e *Equation) String() string {

	return fmt.Sprintf("Equation[%d: %s]", e.TestValue, strings.Join(
		utils.SliceMap(e.Numbers, func(intVal int) string { return strconv.Itoa(intVal) }),
		" ",
	))
}

func (e *Equation) Equal(other Equation) bool {

	return e.TestValue == other.TestValue && utils.SliceEqual(e.Numbers, other.Numbers)
}

// Parse from something like "190: 10 19".
//
// Panic if string can not be parsed.
func MustParseFromLine(line string) *Equation {

	equation, err := TryParseFromLine(line)
	utils.PanicIf(err)

	return equation
}

// Parse from something like "190: 10 19".
func TryParseFromLine(line string) (*Equation, error) {

	tokens := strings.Split(line, ": ")
	if len(tokens) != 2 {
		return nil, fmt.Errorf("expected two tokens separated by ': ' in equation line '%s', got %d", line, len(tokens))
	}

	testValue, err := strconv.Atoi(tokens[0])
	if err != nil {
		return nil, fmt.Errorf("can not parse test value '%s'", tokens[0])
	}

	numberTokens := strings.Split(tokens[1], " ")
	numbers := make([]int, 0, len(numberTokens))

	for _, numberToken := range numberTokens {
		number, err := strconv.Atoi(numberToken)
		if err != nil {
			return nil, fmt.Errorf("can not parse number '%s'", numberToken)
		}
		numbers = append(numbers, number)
	}

	return &Equation{
		TestValue: testValue,
		Numbers:   numbers,
	}, nil
}
