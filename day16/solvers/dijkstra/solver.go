package dijkstra

import (
	"math"

	"github.com/nlduy0310/aoc-2024/day16/config"
	"github.com/nlduy0310/aoc-2024/day16/direction"
	"github.com/nlduy0310/aoc-2024/day16/maze"
	"github.com/nlduy0310/aoc-2024/day16/position"
	"github.com/nlduy0310/aoc-2024/day16/priorityqueue"
	"github.com/nlduy0310/aoc-2024/day16/state"
	"github.com/nlduy0310/aoc-2024/day16/statecost"
)

type DijkstraSolver struct {
	maze      maze.Maze
	traceback map[state.State][]state.State
	dist      map[state.State]float64
	pq        priorityqueue.PriorityQueue
	config    config.Config
}

func NewSolver(maze maze.Maze, config config.Config) *DijkstraSolver {
	return &DijkstraSolver{
		maze:      maze,
		traceback: make(map[state.State][]state.State),
		dist:      make(map[state.State]float64),
		pq:        *priorityqueue.NewEmptyQueue(),
		config:    config,
	}
}

func (solver *DijkstraSolver) expand(sourceState state.State, sourceBestCost float64) {
	nextPossibleStatesCost := []statecost.StateCost{
		statecost.New(sourceState.MovedForward(), solver.config.MoveCost),
		statecost.New(sourceState.TurnedLeft(), solver.config.TurnCost),
		statecost.New(sourceState.TurnedRight(), solver.config.TurnCost),
	}

	for _, possibleStateCost := range nextPossibleStatesCost {
		if !(solver.maze.Contains(possibleStateCost.State.Position) && !solver.maze.IsBlockedAt(possibleStateCost.State.Position)) {
			continue
		}

		currentCost, ok := solver.dist[possibleStateCost.State]
		if !ok || ok && sourceBestCost+possibleStateCost.Cost < currentCost {
			solver.traceback[possibleStateCost.State] = []state.State{sourceState}
			solver.dist[possibleStateCost.State] = sourceBestCost + possibleStateCost.Cost
			solver.pq.Push(statecost.New(possibleStateCost.State, sourceBestCost+possibleStateCost.Cost))
		} else if ok && sourceBestCost+possibleStateCost.Cost == currentCost {
			solver.traceback[possibleStateCost.State] = append(solver.traceback[possibleStateCost.State], sourceState)
		}
	}
}

func (solver *DijkstraSolver) Solve() float64 {
	solver.traceback = make(map[state.State][]state.State)
	solver.dist = make(map[state.State]float64)
	solver.pq = *priorityqueue.NewEmptyQueue()

	initialState := state.NewState(solver.maze.StartPosition, solver.config.InitialDirection)
	solver.dist[initialState] = 0
	solver.pq.Push(statecost.New(initialState, 0))

	halting := false
	haltThreshold := -1.0
	var currentExpandNode *statecost.StateCost = solver.pq.Pop()
	for ; currentExpandNode != nil; currentExpandNode = solver.pq.Pop() {
		if halting && currentExpandNode.Cost > haltThreshold {
			break
		}
		if solver.config.ReachedFinishFunc(solver.maze, currentExpandNode.State) {
			halting = true
			haltThreshold = currentExpandNode.Cost
		}
		solver.expand(currentExpandNode.State, currentExpandNode.Cost)
	}

	endStates, dist := solver.findEndStates()
	if len(endStates) == 0 {
		panic("no paths found")
	}

	return dist
}

func (solver *DijkstraSolver) findEndStates() ([]state.State, float64) {
	states := make([]state.State, 0)
	minDist := math.Inf(+1)

	for _, dir := range direction.Directions() {
		s := state.NewState(solver.maze.EndPosition, dir)
		if c, ok := solver.dist[s]; ok {
			if c <= minDist {
				if c < minDist {
					states = states[:0]
					minDist = c
				}
				states = append(states, s)
			}
		}
	}

	return states, minDist
}

func (solver *DijkstraSolver) Traceback() [][]state.State {
	ret := make([][]state.State, 0)
	endStates, _ := solver.findEndStates()

	for _, endState := range endStates {
		subTraceback := solver.tracebackRecursively(endState)
		ret = append(ret, subTraceback...)
	}

	return ret
}

func (solver *DijkstraSolver) tracebackRecursively(from state.State) [][]state.State {
	priorStates, ok := solver.traceback[from]
	if !ok {
		return [][]state.State{{from}}
	}

	ret := make([][]state.State, 0)
	for _, priorState := range priorStates {
		subTracebacks := solver.tracebackRecursively(priorState)
		for _, subTraceback := range subTracebacks {
			ret = append(ret, append([]state.State{from}, subTraceback...))
		}
	}

	if len(ret) == 0 {
		return [][]state.State{{from}}
	}

	return ret
}

func (solver *DijkstraSolver) CountTiles() int {
	tiles := make(map[position.Position]struct{})

	for _, traceback := range solver.Traceback() {
		for _, state := range traceback {
			if !solver.maze.Contains(state.Position) || solver.maze.IsBlockedAt(state.Position) {
				continue
			}
			tiles[state.Position] = struct{}{}
		}
	}

	return len(tiles)
}
