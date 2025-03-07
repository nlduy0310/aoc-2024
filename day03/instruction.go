package main

type Instruction interface {
	getIndex() int
	isCompleted() bool
}

type DoInstruction struct {
	index       int
	instruction string
}

func NewDoInstruction(idx int) DoInstruction {
	return DoInstruction{
		index:       idx,
		instruction: "do()",
	}
}

func (i DoInstruction) getIndex() int {
	return i.index
}

func (i DoInstruction) isCompleted() bool {
	return true
}

type DontInstruction struct {
	index       int
	instruction string
}

func NewDontInstruction(idx int) DontInstruction {
	return DontInstruction{
		index:       idx,
		instruction: "don't()",
	}
}

func (i DontInstruction) getIndex() int {
	return i.index
}

func (i DontInstruction) isCompleted() bool {
	return true
}

type MulInstruction struct {
	index                     int
	instruction               string
	firstNumber, secondNumber int
	isValid                   bool
}

func NewMulInstruction(idx int) MulInstruction {
	return MulInstruction{
		index:        idx,
		instruction:  "mul(",
		firstNumber:  -1,
		secondNumber: -1,
		isValid:      false,
	}
}

func (i *MulInstruction) TryParse(str string, fromIdx int) bool {
	if !(str[fromIdx:fromIdx+len("mul(")] == "mul(") {
		return false
	}

	firstNumber, ok := findDigits(str[fromIdx+len("mul("):])
	if !ok || len(firstNumber) > maxDigitsCount {
		return false
	}
	i.firstNumber = MustAtoI(firstNumber)

	commaIndex := fromIdx + len("mul(") + len(firstNumber)
	if commaIndex >= len(str) || str[commaIndex] != ',' {
		return false
	}

	secondNumber, ok := findDigits(str[commaIndex+1:])
	if !ok || len(secondNumber) > maxDigitsCount {
		return false
	}
	i.secondNumber = MustAtoI(secondNumber)

	closingBrackerIndex := commaIndex + 1 + len(secondNumber)
	if closingBrackerIndex >= len(str) || str[closingBrackerIndex] != ')' {
		return false
	}
	i.isValid = true
	return true
}

func (i MulInstruction) getOutput() int {
	if !i.isValid {
		panic(0)
	}
	return i.firstNumber * i.secondNumber
}

func (i MulInstruction) getIndex() int {
	return i.index
}

func (i MulInstruction) isCompleted() bool {
	return i.isValid
}
