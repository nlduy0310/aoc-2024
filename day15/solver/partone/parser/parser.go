package parser

import (
	"fmt"

	"github.com/nlduy0310/aoc-2024/day15/direction"
	"github.com/nlduy0310/aoc-2024/day15/solver/partone/mapstate"
	"github.com/nlduy0310/aoc-2024/utils"
)

func TryParseFromFile(file string) (*mapstate.MapState, []direction.Direction, error) {

	lines := utils.MustReadLines(file)
	if len(lines) == 0 {
		return nil, nil, fmt.Errorf("empty input file: '%s'", file)
	}

	splitIdx := -1
	for idx, line := range lines {
		if len(line) == 0 {
			splitIdx = idx
			break
		}
	}

	if splitIdx == -1 {
		return nil, nil, fmt.Errorf("can not find the empty split line between map and instruction")
	}

	mapState, err := mapstate.TryParseFromLines(lines[0:splitIdx])
	if err != nil {
		return nil, nil, fmt.Errorf("can not parse input file '%s': %s", file, err.Error())
	}

	directions := make([]direction.Direction, 0)
	for lineIdx := splitIdx + 1; lineIdx < len(lines); lineIdx++ {
		for _, r := range lines[lineIdx] {
			d, err := direction.TryParseFromRune(r)
			if err != nil {
				return nil, nil, fmt.Errorf("can not parse input file '%s' at line %d: %s", file, lineIdx, err.Error())
			}
			directions = append(directions, d)
		}
	}

	return mapState, directions, nil
}
