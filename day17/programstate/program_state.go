package programstate

import (
	"fmt"

	"github.com/nlduy0310/aoc-2024/day17/register"
)

type ProgramState struct {
	registers map[string]*register.Register
	nextJump  int
	output    []string
}

func Init(registers []*register.Register) *ProgramState {
	registersMap := make(map[string]*register.Register)
	for _, register := range registers {
		registersMap[register.Name] = register
	}

	return &ProgramState{
		registers: registersMap,
		nextJump:  0,
		output:    make([]string, 0),
	}
}

func (s *ProgramState) GetRegisterLiteral(registerName string) (int, error) {
	register, ok := s.registers[registerName]
	if !ok {
		return 0, fmt.Errorf("register \"%s\" not found", registerName)
	}

	return register.LiteralValue, nil
}

func (s *ProgramState) SetRegisterLiteral(registerName string, literalValue int) error {
	register, ok := s.registers[registerName]
	if !ok {
		return fmt.Errorf("register \"%s\" not found", registerName)
	}

	register.LiteralValue = literalValue
	return nil
}

func (s *ProgramState) AddToOutput(str string) {
	s.output = append(s.output, str)
}

func (s *ProgramState) SetJump(jumpVal int) {
	s.nextJump = jumpVal
}

func (s *ProgramState) GetJump() int {
	return s.nextJump
}

func (s *ProgramState) ShiftJump(shiftVal int) {
	s.nextJump += shiftVal
}

func (s *ProgramState) GetOutput() []string {
	return s.output
}
