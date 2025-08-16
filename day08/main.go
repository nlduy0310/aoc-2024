package day08

import (
	"fmt"

	partone_solver "github.com/nlduy0310/aoc-2024/day08/solver/part_one"
	parttwo_solver "github.com/nlduy0310/aoc-2024/day08/solver/part_two"
)

func SolvePartOne(inputFile string) {

	solver := partone_solver.MustInitPartOneSolver(inputFile)
	result := solver.Solve()

	fmt.Printf("Part one: %d\n", result)
}

func SolvePartTwo(inputFile string) {

	solver := parttwo_solver.MustInitPartTwoSolver(inputFile)
	result := solver.Solve()

	fmt.Printf("Part two: %d\n", result)
}
