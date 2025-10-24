package instruction

type Instruction int

const (
	ADV Instruction = iota
	BXL
	BST
	JNZ
	BXC
	OUT
	BDV
	CDV
)

var intMap = map[int]Instruction{
	0: ADV,
	1: BXL,
	2: BST,
	3: JNZ,
	4: BXC,
	5: OUT,
	6: BDV,
	7: CDV,
}

func FromInt(opcode int) (Instruction, error) {
	ret, ok := intMap[opcode]
	if !ok {
		return ADV, generateInvalidInstructionError(Instruction(opcode))
	}

	return ret, nil
}
