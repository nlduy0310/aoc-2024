package program

import (
	"log"
	"strings"

	"github.com/nlduy0310/aoc-2024/day17/instruction"
	"github.com/nlduy0310/aoc-2024/day17/operand"
	"github.com/nlduy0310/aoc-2024/day17/programstate"
)

type Program struct {
	commandSequence    []int
	instructionPointer int
	programState       *programstate.ProgramState
}

func (p *Program) executeOnce() {
	opcode := p.commandSequence[p.instructionPointer]
	ins, err := instruction.FromInt(opcode)
	if err != nil {
		log.Fatalf("invalid opcode: %d", opcode)
	}
	operand := operand.NewOperand(p.commandSequence[p.instructionPointer+1])

	p.programState.SetJump(p.instructionPointer + 2)
	ins.Execute(p.programState, operand)
	p.jump(p.programState.GetJump())
}

func (p *Program) jump(jumpPointer int) {
	p.instructionPointer = jumpPointer
}

func (p *Program) Execute() {
	for p.instructionPointer < len(p.commandSequence)-1 {
		p.executeOnce()
	}
}

func (p *Program) GetOutput() string {
	return strings.Join(p.programState.GetOutput(), ",")
}
