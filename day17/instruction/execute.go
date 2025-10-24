package instruction

import (
	"github.com/nlduy0310/aoc-2024/day17/operand"
	"github.com/nlduy0310/aoc-2024/day17/programstate"
)

type executeFunc = func(*programstate.ProgramState, operand.Operand)

var executeMap = map[Instruction]executeFunc{
	ADV: handleADVInstruction,
	BXL: handleBXLInstruction,
	BST: handleBSTInstruction,
	JNZ: handleJNZInstruction,
	BXC: handleBXCInstruction,
	OUT: handleOUTInstruction,
	BDV: handleBDVInstruction,
	CDV: handleCDVInstruction,
}

func (i Instruction) Execute(s *programstate.ProgramState, operand operand.Operand) {
	executor, ok := executeMap[i]
	if !ok {
		panic(generateInvalidInstructionError(i))
	}

	executor(s, operand)
}
