package instruction

import (
	"github.com/nlduy0310/aoc-2024/day17/operand"
	"github.com/nlduy0310/aoc-2024/day17/programstate"
)

func handleBSTInstruction(s *programstate.ProgramState, operand operand.Operand) {
	val := operand.EvaluateCombo(s) % 8
	mustSetRegisterLiteral(s, "B", val)
}
