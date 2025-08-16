package solvers

import (
	"github.com/nlduy0310/aoc-2024/day07/equation"
	"github.com/nlduy0310/aoc-2024/day07/operator"
	"github.com/nlduy0310/aoc-2024/utils"
)

type PartTwoSolver struct {
	inputFile        string
	equations        []*equation.Equation
	allowedOperators []operator.Operator
}

func MustInitPartTwoSolver(inputFile string) *PartTwoSolver {

	lines := utils.MustReadLines(inputFile)

	equations := utils.SliceMap(lines, func(line string) *equation.Equation {
		return equation.MustParseFromLine(line)
	})

	return &PartTwoSolver{
		inputFile:        inputFile,
		equations:        equations,
		allowedOperators: []operator.Operator{operator.OpAdd, operator.OpMultiply, operator.OpConcatenate},
	}
}

func (solver *PartTwoSolver) Solve() int {

	ret := 0

	for _, equation := range solver.equations {
		if solver.trySolveEquation(equation) {
			ret += equation.TestValue
		}
	}

	return ret
}

func (solver *PartTwoSolver) trySolveEquation(eq *equation.Equation) bool {

	initialTestValue := eq.Numbers[0]

	return solver.solveRecursively(eq, 1, initialTestValue)
}

func (solver *PartTwoSolver) solveRecursively(eq *equation.Equation, valuesGrouped int, currentTestValue int) bool {

	utils.Assert(valuesGrouped > 0, "this function expects the input to be a group of one or more values")

	if valuesGrouped == len(eq.Numbers) {
		return currentTestValue == eq.TestValue
	}

	nextValueIdx := valuesGrouped
	for _, operator := range solver.allowedOperators {
		if ok := solver.solveRecursively(eq, valuesGrouped+1, operator.Apply(currentTestValue, eq.Numbers[nextValueIdx])); ok {
			return true
		}
	}

	return false
}
