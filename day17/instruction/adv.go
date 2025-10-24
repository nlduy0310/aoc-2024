package instruction

import (
	"math"

	"github.com/nlduy0310/aoc-2024/day17/operand"
	"github.com/nlduy0310/aoc-2024/day17/programstate"
)

func handleADVInstruction(s *programstate.ProgramState, operand operand.Operand) {
	numerator := mustGetRegisterLiteral(s, "A")
	operandVal := operand.EvaluateCombo(s)
	val := int(math.Trunc(float64(numerator) / math.Pow(float64(2), float64(operandVal))))
	mustSetRegisterLiteral(s, "A", val)
}
