package instruction

import (
	"github.com/nlduy0310/aoc-2024/day17/operand"
	"github.com/nlduy0310/aoc-2024/day17/programstate"
)

func handleJNZInstruction(s *programstate.ProgramState, operand operand.Operand) {
	registerA := mustGetRegisterLiteral(s, "A")
	if registerA == 0 {
		return
	}

	operandVal := operand.EvaluateLiteral()
	s.SetJump(operandVal)
}
