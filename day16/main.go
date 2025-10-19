package day16

import (
	"fmt"

	"github.com/nlduy0310/aoc-2024/day16/config"
	"github.com/nlduy0310/aoc-2024/day16/direction"
	"github.com/nlduy0310/aoc-2024/day16/maze"
	"github.com/nlduy0310/aoc-2024/day16/solvers/dijkstra"
	"github.com/nlduy0310/aoc-2024/day16/state"
	"github.com/nlduy0310/aoc-2024/utils"
)

func SolvePartOne(file string) {
	m, err := maze.ParseFromFile(file)
	utils.PanicIf(err)

	config := config.New(1, 1000, direction.East, func(m maze.Maze, s state.State) bool {
		return m.EndPosition == s.Position
	})

	solver := dijkstra.NewSolver(*m, config)
	res := solver.Solve()

	fmt.Printf("Part one: %f\n", res)
}

func SolvePartTwo(file string) {
	m, err := maze.ParseFromFile(file)
	utils.PanicIf(err)

	config := config.New(1, 1000, direction.East, func(m maze.Maze, s state.State) bool {
		return m.EndPosition == s.Position
	})

	solver := dijkstra.NewSolver(*m, config)
	solver.Solve()
	res := solver.CountTiles()

	fmt.Printf("Part two: %d\n", res)
}
