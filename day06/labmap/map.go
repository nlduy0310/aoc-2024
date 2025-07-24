package labmap

import (
	"fmt"
	"strings"

	"github.com/nlduy0310/aoc-2024/day06/position"
	"github.com/nlduy0310/aoc-2024/utils"
)

type Map struct {
	cells      [][]cellState
	rows, cols int
}

func (m *Map) String() string {

	var mapContentBuilder strings.Builder

	for _, rowContent := range m.cells {
		for _, cell := range rowContent {
			mapContentBuilder.WriteString(cell.String())
		}
		mapContentBuilder.WriteString("\n")
	}

	mapContent := mapContentBuilder.String()

	return fmt.Sprintf("Map[\nrows=%d\ncols=%d\nmap:\n%s]", m.rows, m.cols, mapContent)
}

func (m *Map) Copy() *Map {

	copiedCells := make([][]cellState, len(m.cells))
	for idx := range copiedCells {
		copiedCells[idx] = utils.SliceCopy(m.cells[idx])
	}

	return &Map{
		cells: copiedCells,
		rows:  m.rows,
		cols:  m.cols,
	}
}

func (m *Map) IsBlocked(position position.Position) bool {

	positionState := m.cells[position.Row][position.Col]

	switch positionState {
	case cellEmpty:
		return false
	case cellBlocked:
		return true
	default:
		panic(fmt.Sprintf("invalid cell state enum: %d", positionState))
	}
}

func NewMapFromLines(lines []string) (*Map, error) {

	if lines == nil {
		return nil, fmt.Errorf("can not construct map from nil value")
	}

	if len(lines) == 0 {
		return nil, fmt.Errorf("can not construct map from empty lines")
	}

	var rows, cols = len(lines), len(lines[0])

	cells := make([][]cellState, rows)

	for r := range rows {

		line := lines[r]

		if len(line) != cols {
			return nil, fmt.Errorf("failed to initialize map from lines. inconsistent row lengths: expected %d, got %d", cols, len(line))
		}

		currentLineCells := make([]cellState, cols)

		for c, cellRune := range line {
			cellValue := string(cellRune)
			switch cellValue {
			case ".":
				currentLineCells[c] = cellEmpty
			case "#":
				currentLineCells[c] = cellBlocked
			case "^":
				currentLineCells[c] = cellEmpty
			default:
				return nil, fmt.Errorf("encountered unexpected value when initializing the map from string: '%s'", cellValue)
			}
		}

		cells[r] = currentLineCells
	}

	return &Map{
		cells: cells,
		rows:  rows,
		cols:  cols,
	}, nil
}

func (m *Map) Rows() int {

	return m.rows
}

func (m *Map) Cols() int {

	return m.cols
}

func (m *Map) Block(row, col int) {

	m.cells[row][col] = cellBlocked
}

func (m *Map) Unblock(row, col int) {

	m.cells[row][col] = cellEmpty
}
