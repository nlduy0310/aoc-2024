package operand

import (
	"fmt"

	"github.com/nlduy0310/aoc-2024/day17/programstate"
)

type Operand struct {
	literalValue int
}

func NewOperand(literalValue int) Operand {
	return Operand{
		literalValue: literalValue,
	}
}

func (o Operand) EvaluateLiteral() int {
	return o.literalValue
}

func mustGetRegisterLiteral(s *programstate.ProgramState, registerName string) int {
	val, err := s.GetRegisterLiteral(registerName)
	if err != nil {
		panic(fmt.Sprintf("can not get literal from register \"%s\": %s", registerName, err.Error()))
	}

	return val
}

func (o Operand) EvaluateCombo(s *programstate.ProgramState) int {
	switch o.literalValue {
	case 0, 1, 2, 3:
		return o.literalValue
	case 4:
		return mustGetRegisterLiteral(s, "A")
	case 5:
		return mustGetRegisterLiteral(s, "B")
	case 6:
		return mustGetRegisterLiteral(s, "C")
	default:
		panic(fmt.Sprintf("unsupported combo operand %d", o.literalValue))
	}
}
