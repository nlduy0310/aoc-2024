package config

import (
	"github.com/nlduy0310/aoc-2024/day16/direction"
	"github.com/nlduy0310/aoc-2024/day16/maze"
	"github.com/nlduy0310/aoc-2024/day16/state"
)

type Config struct {
	MoveCost          float64
	TurnCost          float64
	InitialDirection  direction.Direction
	ReachedFinishFunc func(maze.Maze, state.State) bool
}

func New(moveCost, turnCost float64, initialDirection direction.Direction, reachedFinishFunc func(maze.Maze, state.State) bool) Config {
	return Config{
		MoveCost:          moveCost,
		TurnCost:          turnCost,
		InitialDirection:  initialDirection,
		ReachedFinishFunc: reachedFinishFunc,
	}
}
