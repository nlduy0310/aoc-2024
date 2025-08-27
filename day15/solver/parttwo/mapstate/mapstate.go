package mapstate

import (
	"fmt"
	"strings"

	"github.com/nlduy0310/aoc-2024/day15/position"
	"github.com/nlduy0310/aoc-2024/day15/solver/parttwo/box"
	"github.com/nlduy0310/aoc-2024/utils"
)

type MapState struct {
	Width, Height int
	Walls         map[position.Position]struct{}
	Boxes         []*box.Box
	Robot         position.Position
}

func TryParseFromLines(lines []string) (*MapState, error) {

	if len(lines) == 0 {
		return nil, fmt.Errorf("no lines provided")
	}

	width, height := 2*len(lines[0]), len(lines)
	boxes := make([]*box.Box, 0)
	walls := make(map[position.Position]struct{})
	var robot *position.Position = nil

	for lineIdx, line := range lines {
		if 2*len(line) != width {
			return nil, fmt.Errorf("inconsistent line length at line 0 (%d) and line %d (%d)", width/2, lineIdx, len(line))
		}

		for runeIdx, r := range line {
			switch r {
			case '#':
				walls[position.NewPosition(lineIdx, 2*runeIdx)] = struct{}{}
				walls[position.NewPosition(lineIdx, 2*runeIdx+1)] = struct{}{}
			case 'O':
				newBox := box.NewBox(
					position.NewPosition(lineIdx, 2*runeIdx),
					position.NewPosition(lineIdx, 2*runeIdx+1),
				)
				boxes = append(boxes, &newBox)
			case '@':
				p := position.NewPosition(lineIdx, 2*runeIdx)
				if robot != nil {
					return nil, fmt.Errorf("multiple robots found at %s and %s", robot.String(), p.String())
				}
				robot = &p
			default:
				continue
			}
		}
	}

	if robot == nil {
		return nil, fmt.Errorf("no robots found")
	}

	return &MapState{
		Width:  width,
		Height: height,
		Boxes:  boxes,
		Walls:  walls,
		Robot:  *robot,
	}, nil
}

func (m *MapState) ContainsPosition(p position.Position) bool {

	return utils.IsInRangeInclusive(p.Row, 0, m.Height-1) &&
		utils.IsInRangeInclusive(p.Col, 0, m.Width-1)
}

func (m *MapState) IsWall(p position.Position) bool {

	if !m.ContainsPosition(p) {
		return false
	}

	_, ok := m.Walls[p]

	return ok
}

func (m *MapState) IsEmpty(p position.Position) bool {

	if !m.ContainsPosition(p) {
		return false
	}

	if m.IsWall(p) {
		return false
	}

	if _, err := m.TryGetBoxAt(p); err == nil {
		return false
	}

	return true
}

func (m *MapState) TryGetBoxAt(p position.Position) (*box.Box, error) {

	notFoundErr := fmt.Errorf("box not found at %s", p.String())

	if !m.ContainsPosition(p) {
		return nil, notFoundErr
	}

	for _, box := range m.Boxes {
		if box.ContainsPosition(p) {
			return box, nil
		}
	}

	return nil, notFoundErr
}

func (m *MapState) Visualize() {

	mat := make([][]int, m.Height)
	for i := range mat {
		mat[i] = make([]int, m.Width)
	}

	for wall := range m.Walls {
		mat[wall.Row][wall.Col] = 1
	}

	for _, b := range m.Boxes {
		mat[b.LeftPosition.Row][b.LeftPosition.Col] = 2
		mat[b.RightPosition.Row][b.RightPosition.Col] = 3
	}

	mat[m.Robot.Row][m.Robot.Col] = 4

	builder := strings.Builder{}

	for _, row := range mat {
		for _, cell := range row {
			switch cell {
			case 1:
				builder.WriteRune('#')
			case 2:
				builder.WriteRune('[')
			case 3:
				builder.WriteRune(']')
			case 4:
				builder.WriteRune('@')
			default:
				builder.WriteRune('.')
			}
		}
		builder.WriteRune('\n')
	}

	print(builder.String())
}
