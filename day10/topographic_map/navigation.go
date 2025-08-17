package topographicmap

import (
	"github.com/nlduy0310/aoc-2024/day10/position"
	"github.com/nlduy0310/aoc-2024/utils"
)

func (m *TopographicMap) GetPossibleTrailheads() []position.Position {

	ret := make([]position.Position, 0)

	for rowIdx, row := range m.cells {
		for colIdx, cellValue := range row {
			if cellValue == 0 {
				ret = append(ret, position.Position{Row: rowIdx, Col: colIdx})
			}
		}
	}

	return ret
}

func (m *TopographicMap) GetValueAt(position position.Position) int {

	return m.cells[position.Row][position.Col]
}

func (m *TopographicMap) IsWithinMap(position position.Position) bool {

	return utils.IsInRangeInclusive(position.Row, 0, m.height-1) &&
		utils.IsInRangeInclusive(position.Col, 0, m.width-1)
}

func (m *TopographicMap) GetNearbyCells(targetPosition position.Position) []position.Position {

	ret := make([]position.Position, 0)

	for _, nearbyCell := range []position.Position{targetPosition.Up(), targetPosition.Down(), targetPosition.Left(), targetPosition.Right()} {
		if m.IsWithinMap(nearbyCell) {
			ret = append(ret, nearbyCell)
		}
	}

	return ret
}
