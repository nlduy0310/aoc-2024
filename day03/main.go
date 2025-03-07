package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

const inputFilePath = "./day03/input"
const maxDigitsCount = 3

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func getInput(filePath string) string {
	data, err := os.ReadFile(filePath)
	check(err)
	return string(data)
}

func findPrefixIndexes(str string, prfx string) []int {
	res := make([]int, 0)
	findFrom := 0
	for {
		tmp := strings.Index(str[findFrom:], prfx)
		if tmp == -1 {
			break
		}
		res = append(res, findFrom+tmp)
		findFrom += tmp + len(prfx)
	}
	return res
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func MustRtoI(b byte) int {
	if !isDigit(b) {
		panic(0)
	}
	return int(b - '0')
}

func MustAtoI(str string) int {
	res, err := strconv.Atoi(str)
	check(err)
	return res
}

func findDigits(str string) (string, bool) {
	if !isDigit(str[0]) {
		return "", false
	}

	i := 0
	for i+1 < len(str) && isDigit(str[i+1]) {
		i++
	}
	return str[0 : i+1], true
}

func main() {
	input := getInput(inputFilePath)

	firstResult := partOne(input)
	secondResult := partTwo(input)
	fmt.Printf("Part one: %d\n", firstResult)
	fmt.Printf("Part two: %d\n", secondResult)
}

func partOne(input string) int {

	potentialIndexes := findPrefixIndexes(input, "mul(")
	res := 0

	for _, potentialIndex := range potentialIndexes {
		firstNumber, ok := findDigits(input[potentialIndex+len("mul("):])
		if !ok || len(firstNumber) > 3 {
			continue
		}
		commaIndex := potentialIndex + len("mul(") + len(firstNumber)
		if commaIndex >= len(input) || input[commaIndex] != ',' {
			continue
		}

		secondNumber, ok := findDigits(input[commaIndex+1:])
		if !ok || len(secondNumber) > 3 {
			continue
		}
		closingBracketIndex := commaIndex + 1 + len(secondNumber)
		if closingBracketIndex >= len(input) || input[closingBracketIndex] != ')' {
			continue
		}

		a, b := MustAtoI(firstNumber), MustAtoI(secondNumber)
		res += a * b
	}

	return res
}

func partTwo(input string) int {
	instructions := make([]Instruction, 0)

	for _, v := range findPrefixIndexes(input, "do()") {
		instructions = append(instructions, NewDoInstruction(v))
	}

	for _, v := range findPrefixIndexes(input, "don't()") {
		instructions = append(instructions, NewDontInstruction(v))
	}

	for _, v := range findPrefixIndexes(input, "mul(") {
		m := NewMulInstruction(v)
		if m.TryParse(input, m.index) {
			instructions = append(instructions, m)
		}
	}

	slices.SortFunc(instructions, func(a, b Instruction) int {
		return a.getIndex() - b.getIndex()
	})

	total := 0
	enabled := true
	for _, ins := range instructions {
		if _, ok := ins.(DoInstruction); ok {
			enabled = true
		} else if _, ok := ins.(DontInstruction); ok {
			enabled = false
		} else if mulIns, ok := ins.(MulInstruction); ok {
			if enabled {
				total += mulIns.getOutput()
			}
		}
	}
	return total
}
