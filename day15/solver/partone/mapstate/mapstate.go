package mapstate

import (
	"fmt"

	"github.com/nlduy0310/aoc-2024/day15/position"
	"github.com/nlduy0310/aoc-2024/utils"
)

type MapState struct {
	Width, Height int
	Walls         []position.Position
	Boxes         []position.Position
	Robot         position.Position
}

func TryParseFromLines(lines []string) (*MapState, error) {

	if len(lines) == 0 {
		return nil, fmt.Errorf("no lines provided")
	}

	width, height := len(lines[0]), len(lines)
	walls := make([]position.Position, 0)
	boxes := make([]position.Position, 0)
	var robot *position.Position = nil

	for lineIdx, line := range lines {
		if len(line) != width {
			return nil, fmt.Errorf("in consistent length at line 0 (%d) and line %d (%d)", width, lineIdx, len(line))
		}

		for runeIdx, r := range line {
			switch r {
			case '#':
				walls = append(walls, position.NewPosition(lineIdx, runeIdx))
			case 'O':
				boxes = append(boxes, position.NewPosition(lineIdx, runeIdx))
			case '@':
				if robot != nil {
					return nil, fmt.Errorf("multiple robots detected, at %s and %s", robot.String(), position.NewPosition(lineIdx, runeIdx).String())
				}
				p := position.NewPosition(lineIdx, runeIdx)
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
		Walls:  walls,
		Boxes:  boxes,
		Robot:  *robot,
	}, nil
}

func MustParseFromLines(lines []string) MapState {

	ret, err := TryParseFromLines(lines)
	utils.ExitIf(err)

	return *ret
}

func (m MapState) ContainsPosition(p position.Position) bool {

	return utils.IsInRangeInclusive(p.Row, 0, m.Height-1) &&
		utils.IsInRangeInclusive(p.Col, 0, m.Width-1)
}

func (m MapState) IsBox(p position.Position) bool {

	if !m.ContainsPosition(p) {
		return false
	}

	for _, box := range m.Boxes {
		if p == box {
			return true
		}
	}

	return false
}

func (m MapState) IsWall(p position.Position) bool {

	if !m.ContainsPosition(p) {
		return false
	}

	for _, wall := range m.Walls {
		if p == wall {
			return true
		}
	}

	return false
}

func (m MapState) TryGetBox(p position.Position) *position.Position {

	if !m.ContainsPosition(p) {
		return nil
	}

	for idx := range m.Boxes {
		if p == m.Boxes[idx] {
			return &m.Boxes[idx]
		}
	}

	return nil
}

func (m MapState) TryGetWall(p position.Position) *position.Position {

	if !m.ContainsPosition(p) {
		return nil
	}

	for idx := range m.Walls {
		if p == m.Walls[idx] {
			return &m.Walls[idx]
		}
	}

	return nil
}
