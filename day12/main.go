package day12

import (
	"fmt"

	"github.com/nlduy0310/aoc-2024/day12/solver/partone"
	"github.com/nlduy0310/aoc-2024/day12/solver/parttwo"
)

func SolvePartOne(inputFile string) {

	solver := partone.MustInitSolver(inputFile)
	price := solver.Solve()

	fmt.Printf("Part one: %d\n", price)
}

func SolvePartTwo(inputFile string) {

	solver := parttwo.MustInitSolver(inputFile)
	price := solver.Solve()

	fmt.Printf("Part two: %d\n", price)
}
