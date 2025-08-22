package garden

import (
	"fmt"

	"github.com/nlduy0310/aoc-2024/day12/position"
	"github.com/nlduy0310/aoc-2024/utils"
)

type Garden struct {
	Width, Height int
	cells         [][]rune
}

func MustParseFromFile(file string) Garden {

	lines := utils.MustReadLines(file)
	utils.Assert(len(lines) > 0, fmt.Sprintf("empty input file '%s'", file))

	width, height := len(lines[0]), len(lines)

	cells := make([][]rune, len(lines))

	for lineIdx, line := range lines {
		utils.Assert(len(line) == width, fmt.Sprintf("inconsistent line lengths in input file '%s': line 0 have length %d but line %d have length %d", file, width, lineIdx, len(line)))

		cellsLine := make([]rune, len(line))

		for charIdx, char := range line {
			cellsLine[charIdx] = char
		}

		cells[lineIdx] = cellsLine
	}

	return Garden{
		Width:  width,
		Height: height,
		cells:  cells,
	}
}

func (g Garden) GetPlantAt(position position.Position) (*rune, error) {

	if !utils.IsInRangeInclusive(position.Row, 0, g.Height-1) {
		return nil, fmt.Errorf("invalid row index %d: out of range [%d, %d]", position.Row, 0, g.Height-1)
	}
	if !utils.IsInRangeInclusive(position.Col, 0, g.Width-1) {
		return nil, fmt.Errorf("invalid col index %d: out of range [%d, %d]", position.Col, 0, g.Width-1)
	}

	return &g.cells[position.Row][position.Col], nil
}

func (g Garden) Contains(position position.Position) bool {

	return utils.IsInRangeInclusive(position.Row, 0, g.Height-1) && utils.IsInRangeInclusive(position.Col, 0, g.Width-1)
}

func (g Garden) GetSurroundingPositions(currentPosition position.Position) []position.Position {

	ret := make([]position.Position, 0)

	for _, surroundingPosition := range []position.Position{currentPosition.Up(), currentPosition.Down(), currentPosition.Left(), currentPosition.Right()} {
		if g.Contains(surroundingPosition) {
			ret = append(ret, surroundingPosition)
		}
	}

	return ret
}

func (g Garden) GetSurroundingPositionsFiltered(currentPosition position.Position, filter func(position.Position) bool) []position.Position {

	ret := make([]position.Position, 0)

	for _, surroundingPosition := range []position.Position{currentPosition.Up(), currentPosition.Down(), currentPosition.Left(), currentPosition.Right()} {
		if g.Contains(surroundingPosition) && filter(surroundingPosition) {
			ret = append(ret, surroundingPosition)
		}
	}

	return ret
}
