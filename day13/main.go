package day13

import (
	"fmt"

	"github.com/nlduy0310/aoc-2024/day13/solver/partone"
	"github.com/nlduy0310/aoc-2024/day13/solver/parttwo"
)

func SolvePartOne(inputFile string) {

	solver := partone.MustInitSolver(inputFile)
	result := solver.Solve()

	fmt.Printf("Part one: %d\n", result)
}

func SolvePartTwo(inputFile string) {

	solver := parttwo.MustInitSolver(inputFile)
	result := solver.Solve()

	fmt.Printf("Part two: %d\n", result)
}
