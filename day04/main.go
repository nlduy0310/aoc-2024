package main

import (
	"fmt"
)

const (
	PART_ONE_INPUT = "day04/data/part-one"
	PART_TWO_INPUT = "day04/data/part-two"
)

func main() {

	partOneRunesBoard, err := readInputFile(PART_ONE_INPUT)
	panicIf(err, fmt.Sprintf("can not read input file '%s'", PART_ONE_INPUT))
	partTwoRunesBoard, err := readInputFile(PART_TWO_INPUT)
	panicIf(err, fmt.Sprintf("can not read input file '%s'", PART_TWO_INPUT))

	partOneResult := solvePartOne(partOneRunesBoard)
	partTwoResult := solvePartTwo(partTwoRunesBoard)

	fmt.Printf("Part one: %d\n", partOneResult)
	fmt.Printf("Part two: %d\n", partTwoResult)
}
