package region

import (
	"slices"

	"github.com/nlduy0310/aoc-2024/day12/direction"
	"github.com/nlduy0310/aoc-2024/day12/position"
)

type Region struct {
	Plant     rune
	positions map[position.Position]struct{}
}

func NewRegion(plant rune) Region {

	return Region{Plant: plant, positions: make(map[position.Position]struct{})}
}

func (r Region) AddPosition(position position.Position) {

	r.positions[position] = struct{}{}
}

func (r Region) Contains(p position.Position) bool {

	_, ok := r.positions[p]
	return ok
}

func (r Region) GetArea() int {

	return len(r.positions)
}

func (r Region) GetPerimeter() int {

	ret := 0

	for regionPosition := range r.positions {
		for _, nearbyPosition := range []position.Position{regionPosition.Up(), regionPosition.Down(), regionPosition.Left(), regionPosition.Right()} {
			if _, ok := r.positions[nearbyPosition]; !ok {
				ret += 1
			}
		}
	}

	return ret
}

func (r Region) GetSidesCount() int {

	sidesMap := make(map[direction.Direction]map[int][]int)
	for _, d := range []direction.Direction{direction.North, direction.South, direction.East, direction.West} {
		sidesMap[d] = make(map[int][]int)
	}

	inspectDirection := func(p position.Position, d direction.Direction) {
		outer := p.Move(d)
		if !r.Contains(outer) {
			var mainAxis, crossAxis int
			if d == direction.North || d == direction.South {
				mainAxis, crossAxis = p.Row, p.Col
			} else {
				mainAxis, crossAxis = p.Col, p.Row
			}

			records, ok := sidesMap[d][mainAxis]
			if !ok {
				sidesMap[d][mainAxis] = []int{crossAxis}
			} else if !slices.Contains(records, crossAxis) {
				sidesMap[d][mainAxis] = append(records, crossAxis)
			}
		}
	}

	for _, regionPosition := range r.GetPositions() {
		inspectDirection(regionPosition, direction.North)
		inspectDirection(regionPosition, direction.South)
		inspectDirection(regionPosition, direction.East)
		inspectDirection(regionPosition, direction.West)
	}

	ret := 0

	for _, d := range []direction.Direction{direction.North, direction.South, direction.East, direction.West} {
		directionMap := sidesMap[d]
		for _, records := range directionMap {
			slices.Sort(records)
			groups := 1
			for i := 0; i < len(records)-1; i++ {
				if records[i+1]-records[i] > 1 {
					groups++
				}
			}
			ret += groups
		}
	}

	return ret
}

func (r Region) GetPositions() []position.Position {

	ret := make([]position.Position, 0, len(r.positions))

	for k := range r.positions {
		ret = append(ret, k)
	}

	return ret
}
