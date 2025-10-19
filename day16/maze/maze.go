package maze

import (
	"fmt"
	"strings"

	"github.com/nlduy0310/aoc-2024/day16/position"
	"github.com/nlduy0310/aoc-2024/utils"
)

type Maze struct {
	StartPosition, EndPosition position.Position
	Width, Height              int
	wallsMap                   map[position.Position]struct{}
}

func NewMaze(startPos, endPos position.Position, width, height int, wallsMap map[position.Position]struct{}) *Maze {
	return &Maze{
		StartPosition: startPos,
		EndPosition:   endPos,
		Width:         width,
		Height:        height,
		wallsMap:      wallsMap,
	}
}

func (m *Maze) Contains(p position.Position) bool {
	return utils.IsInRangeInclusive(p.X, 0, m.Width-1) &&
		utils.IsInRangeInclusive(p.Y, 0, m.Height-1)
}

func (m *Maze) IsBlockedAt(p position.Position) bool {
	_, ok := m.wallsMap[p]
	return ok
}

func (m *Maze) String() string {
	return fmt.Sprintf(
		"Maze[Width=%d, Height=%d, Start=%s, End=%s, Walls=%d]",
		m.Width, m.Height, m.StartPosition, m.EndPosition, len(m.wallsMap),
	)
}

func (m *Maze) PrettyString() string {
	stringBuilder := strings.Builder{}

	stringBuilder.WriteString(m.String() + "\n")
	for lineIdx := range m.Height {
		for colIdx := range m.Width {
			pos := position.NewPosition(colIdx, lineIdx)
			if m.IsBlockedAt(pos) {
				stringBuilder.WriteString(wallBlockChar)
			} else if pos == m.StartPosition {
				stringBuilder.WriteString(startBlockChar)
			} else if pos == m.EndPosition {
				stringBuilder.WriteString(endBlockChar)
			} else {
				stringBuilder.WriteString(emptyBlockChar)
			}
		}
		stringBuilder.WriteString("\n")
	}

	return stringBuilder.String()
}
