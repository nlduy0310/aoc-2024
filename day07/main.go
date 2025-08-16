package day07

import (
	"fmt"

	"github.com/nlduy0310/aoc-2024/day07/solvers"
)

func SolvePartOne(inputFile string) {

	solver := solvers.MustInitPartOneSolver(inputFile)
	res := solver.Solve()

	fmt.Printf("Part one: %d\n", res)
}

func SolvePartTwo(inputFile string) {

	solver := solvers.MustInitPartTwoSolver(inputFile)
	res := solver.Solve()

	fmt.Printf("Part two: %d\n", res)
}
