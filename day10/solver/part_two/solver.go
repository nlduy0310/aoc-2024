package parttwo

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

	ret := 0

	for _, possibleTrailhead := range s.topographicMap.GetPossibleTrailheads() {
		if ok, scores := s.solveTrailhead(possibleTrailhead); ok {
			ret += scores
		}
	}

	return ret
}

func (s *Solver) solveTrailhead(trailhead position.Position) (bool, int) {

	scores := s.solveRecursively(trailhead)

	if scores > 0 {
		return true, scores
	} else {
		return false, 0
	}
}

func (s *Solver) solveRecursively(currentPosition position.Position) int {

	currentValue := s.topographicMap.GetValueAt(currentPosition)
	if currentValue == 9 {
		return 1
	}

	ret := 0
	for _, nearbyCell := range s.topographicMap.GetNearbyCells(currentPosition) {
		if s.topographicMap.GetValueAt(nearbyCell) != currentValue+1 {
			continue
		}

		ret += s.solveRecursively(nearbyCell)
	}

	return ret
}
