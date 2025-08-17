package partone

import (
	"github.com/nlduy0310/aoc-2024/day10/position"
	topographicmap "github.com/nlduy0310/aoc-2024/day10/topographic_map"
)

type Solver struct {
	topographicMap topographicmap.TopographicMap
}

func NewSolver(topoMap topographicmap.TopographicMap) Solver {

	return Solver{
		topographicMap: topoMap,
	}
}

func (s *Solver) Solve() int {

	possibleTrailheads := s.topographicMap.GetPossibleTrailheads()

	ret := 0
	for _, possibleTrailhead := range possibleTrailheads {
		if ok, score := s.solveTrailhead(possibleTrailhead); ok {
			ret += score
		}
	}

	return ret
}

func (s *Solver) solveTrailhead(trailhead position.Position) (bool, int) {

	trailendPositions := make(map[position.Position]struct{})
	onTrailendFound := func(trailend position.Position) {
		trailendPositions[trailend] = struct{}{}
	}
	s.solveRecursively(trailhead, onTrailendFound)

	scores := len(trailendPositions)
	if scores > 0 {
		return true, scores
	} else {
		return false, 0
	}
}

func (s *Solver) solveRecursively(currentPosition position.Position, onTrailendFound func(position.Position)) {

	currentValue := s.topographicMap.GetValueAt(currentPosition)
	if currentValue == 9 {
		onTrailendFound(currentPosition)
		return
	}

	for _, nearbyCell := range s.topographicMap.GetNearbyCells(currentPosition) {
		if s.topographicMap.GetValueAt(nearbyCell) != currentValue+1 {
			continue
		}

		s.solveRecursively(nearbyCell, onTrailendFound)
	}
}
