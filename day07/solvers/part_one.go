package solvers

import (
	"github.com/nlduy0310/aoc-2024/day07/equation"
	"github.com/nlduy0310/aoc-2024/day07/operator"
	"github.com/nlduy0310/aoc-2024/utils"
)

type PartOneSolver struct {
	inputFile        string
	equations        []*equation.Equation
	allowedOperators []operator.Operator
}

type solveState struct {
	stateEquation *equation.Equation
	opsList       []operator.Operator
	currentVal    int
}

func MustInitPartOneSolver(inputFile string) *PartOneSolver {

	lines := utils.MustReadLines(inputFile)

	equations := utils.SliceMap(lines, func(line string) *equation.Equation {
		return equation.MustParseFromLine(line)
	})

	return &PartOneSolver{
		inputFile:        inputFile,
		equations:        equations,
		allowedOperators: []operator.Operator{operator.OpAdd, operator.OpMultiply},
	}
}

func (solver *PartOneSolver) Solve() int {

	ret := 0

	for _, currentEquation := range solver.equations {
		if _, ok := solver.solveEquation(*currentEquation); ok {
			ret += currentEquation.TestValue
		}
	}

	return ret
}

func (solver *PartOneSolver) solveEquation(equation equation.Equation) ([]operator.Operator, bool) {

	if len(equation.Numbers) == 1 {
		return []operator.Operator{}, equation.TestValue == equation.Numbers[0]
	}

	opsList := make([]operator.Operator, 0, len(equation.Numbers)-1)
	initialState := solveState{
		stateEquation: &equation,
		opsList:       opsList,
		currentVal:    0,
	}

	return solver.solveEquationRecursively(&initialState)
}

func (solver *PartOneSolver) solveEquationRecursively(state *solveState) ([]operator.Operator, bool) {

	maxOps := len(state.stateEquation.Numbers) - 1

	if len(state.opsList) == maxOps {
		if state.currentVal != state.stateEquation.TestValue {
			return nil, false
		} else {
			return utils.SliceCopy(state.opsList), true
		}
	}

	for _, allowedOp := range solver.allowedOperators {

		previousVal := state.currentVal
		if len(state.opsList) == 0 {
			state.currentVal = allowedOp.Apply(state.stateEquation.Numbers[0], state.stateEquation.Numbers[1])
		} else {
			nextNumber := state.stateEquation.Numbers[len(state.opsList)+1]
			state.currentVal = allowedOp.Apply(previousVal, nextNumber)
		}
		state.opsList = append(state.opsList, allowedOp)

		ops, ok := solver.solveEquationRecursively(state)
		if ok {
			return ops, true
		}

		state.currentVal = previousVal
		state.opsList = state.opsList[:len(state.opsList)-1]
	}

	return nil, false
}
