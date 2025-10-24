package instruction

import (
	"github.com/nlduy0310/aoc-2024/day17/operand"
	"github.com/nlduy0310/aoc-2024/day17/programstate"
)

func handleBXLInstruction(s *programstate.ProgramState, operand operand.Operand) {
	firstOperand := mustGetRegisterLiteral(s, "B")
	secondOperand := operand.EvaluateLiteral()
	val := firstOperand ^ secondOperand
	mustSetRegisterLiteral(s, "B", val)
}
