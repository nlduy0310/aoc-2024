package instruction

import (
	"fmt"

	"github.com/nlduy0310/aoc-2024/day17/programstate"
)

func generateInvalidInstructionError(val Instruction) error {
	return fmt.Errorf("invalid instruction enum value: %d", val)
}

func mustGetRegisterLiteral(s *programstate.ProgramState, registerName string) int {
	ret, err := s.GetRegisterLiteral(registerName)
	if err != nil {
		panic(fmt.Errorf("can not get literal from register \"%s\": %s", registerName, err.Error()))
	}

	return ret
}

func mustSetRegisterLiteral(s *programstate.ProgramState, registerName string, literalVal int) {
	err := s.SetRegisterLiteral(registerName, literalVal)
	if err != nil {
		panic(fmt.Sprintf("can not set register \"%s\": %s", registerName, err.Error()))
	}
}
