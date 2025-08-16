package day09

import (
	"fmt"

	"github.com/nlduy0310/aoc-2024/day09/diskmap"
	"github.com/nlduy0310/aoc-2024/utils"
)

func SolvePartOne(inputFile string) {

	lines := utils.MustReadLines(inputFile)
	diskMap := diskmap.MustParseFragmentedFromLine(lines[0])
	diskMap.Compact()
	checksum := diskMap.CalculateChecksum()

	fmt.Printf("Part one: %d\n", checksum)
}

func SolvePartTwo(inputFile string) {

	lines := utils.MustReadLines(inputFile)
	diskMap := diskmap.MustParseContiguousFromLine(lines[0])
	diskMap.Compact()
	checksum := diskMap.CalculateChecksum()

	fmt.Printf("Part two: %d\n", checksum)
}
