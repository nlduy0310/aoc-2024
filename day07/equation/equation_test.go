package equation_test

import (
	"testing"

	"github.com/nlduy0310/aoc-2024/day07/equation"
)

func TestEquationParse(t *testing.T) {

	line := "190: 19 10"
	parsedEquation, err := equation.TryParseFromLine(line)
	if err != nil {
		t.Errorf("failed to parse equation from '%s', error: %s", line, err.Error())
	}

	expectedEquation := equation.Equation{
		TestValue: 190,
		Numbers:   []int{19, 10},
	}

	if !parsedEquation.Equal(expectedEquation) {
		t.Errorf("expected %s, got %s", expectedEquation.String(), parsedEquation.String())
	}
}
