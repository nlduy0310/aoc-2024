package instruction

import (
	"strconv"

	"github.com/nlduy0310/aoc-2024/day17/operand"
	"github.com/nlduy0310/aoc-2024/day17/programstate"
)

func handleOUTInstruction(s *programstate.ProgramState, operand operand.Operand) {
	operandVal := operand.EvaluateCombo(s)
	val := operandVal % 8
	s.AddToOutput(strconv.Itoa(val))
}
