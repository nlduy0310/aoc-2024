package day14

import (
	"fmt"

	"github.com/nlduy0310/aoc-2024/day14/solver/partone"
	"github.com/nlduy0310/aoc-2024/day14/solver/parttwo"
)

func SolvePartOne(inputFile string) {

	solver := partone.MustInitSolver(inputFile, 101, 103, 100)
	result := solver.Solve()

	fmt.Printf("Part one: %d\n", result)
}

func SolvePartTwo(inputFile string) {

	solver := parttwo.MustInitSolver(inputFile, 101, 103)
	result := solver.Solve()

	fmt.Printf("Part two: %d\n", result)
}
