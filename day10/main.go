package day10

import (
	"fmt"

	partone "github.com/nlduy0310/aoc-2024/day10/solver/part_one"
	parttwo "github.com/nlduy0310/aoc-2024/day10/solver/part_two"
	topographicmap "github.com/nlduy0310/aoc-2024/day10/topographic_map"
)

func SolvePartOne(inputFile string) {

	topoMap := topographicmap.MustParseFromFile(inputFile)
	solver := partone.NewSolver(topoMap)
	totalScores := solver.Solve()

	fmt.Printf("Part one: %d\n", totalScores)
}

func SolvePartTwo(inputFile string) {

	topoMap := topographicmap.MustParseFromFile(inputFile)
	solver := parttwo.NewSolver(topoMap)
	totalRatings := solver.Solve()

	fmt.Printf("Part two: %d\n", totalRatings)
}
