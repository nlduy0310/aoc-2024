package day06

import (
	"fmt"

	"github.com/nlduy0310/aoc-2024/day06/solvers"
)

var PART_ONE_INPUT string = "./day06/data/input-1"
var PART_TWO_INPUT string = "./day06/data/input-2"

func SolvePartOne() {

	partOneSolver := solvers.MustInitPartOneSolver(PART_ONE_INPUT)
	partOneResult := partOneSolver.Solve()

	fmt.Printf("Part one: %d\n", partOneResult)
}

func SolvePartTwo() {

	partTwoSolver := solvers.MustInitPartTwoSolver(PART_TWO_INPUT)
	partTwoResult := partTwoSolver.Solve()

	fmt.Printf("Part two: %d\n", partTwoResult)
}
