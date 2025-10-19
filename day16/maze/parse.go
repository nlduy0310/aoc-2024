package maze

import (
	"fmt"

	"github.com/nlduy0310/aoc-2024/day16/position"
	"github.com/nlduy0310/aoc-2024/utils"
)

func ParseFromFile(file string) (*Maze, error) {
	lines, err := utils.ReadLines(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read maze file \"%s\": %w", file, err)
	}
	if len(lines) == 0 {
		return nil, fmt.Errorf("can not parse maze from empty file \"%s\"", file)
	}

	var startPos, endPos *position.Position = nil, nil
	width, height := len(lines[0]), len(lines)
	wallsMap := make(map[position.Position]struct{})

	for lineIdx, line := range lines {
		if len(line) != width {
			return nil, fmt.Errorf(
				"failed to read maze file \"%s\": "+
					"invalid line width on line 0 (%d) and line %d (%d)",
				file, width, lineIdx, len(line),
			)
		}

		for colIdx, char := range line {
			pos := position.NewPosition(colIdx, lineIdx)
			switch string(char) {
			case wallBlockChar:
				wallsMap[pos] = struct{}{}
			case startBlockChar:
				if startPos != nil {
					return nil, fmt.Errorf(
						"failed to read maze file \"%s\": "+
							"found 2 start positions at %s and %s",
						file, startPos.String(), pos.String(),
					)
				}
				startPos = &pos
			case endBlockChar:
				if endPos != nil {
					return nil, fmt.Errorf(
						"failed to read maze file \"%s\": "+
							"found 2 end positions at %s and %s",
						file, endPos.String(), pos.String(),
					)
				}
				endPos = &pos
			case emptyBlockChar:
			default:
				return nil, fmt.Errorf(
					"failed to read maze file \"%s\": "+
						"invalid character \"%s\" at %s",
					file, string(char), pos,
				)
			}
		}
	}

	if startPos == nil {
		return nil, fmt.Errorf("can not find start position in maze file \"%s\"", file)
	}
	if endPos == nil {
		return nil, fmt.Errorf("can not find end position in maze file \"%s\"", file)
	}

	return NewMaze(*startPos, *endPos, width, height, wallsMap), nil
}
