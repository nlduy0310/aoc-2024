package parser

import (
	"fmt"
	"slices"

	"github.com/nlduy0310/aoc-2024/day08/antenna"
	"github.com/nlduy0310/aoc-2024/day08/citymap"
	"github.com/nlduy0310/aoc-2024/day08/location"
	"github.com/nlduy0310/aoc-2024/utils"
)

type MapParser struct {
	emptyLocationRunes   []rune
	antennaLocationRunes []rune
}

func NewMapParser(options ...MapParserOption) *MapParser {

	parser := MapParser{
		emptyLocationRunes:   defaultEmptyLocationRunes,
		antennaLocationRunes: defaultAntennaLocationRunes,
	}

	for _, option := range options {
		option(&parser)
	}

	return &parser
}

func (parser *MapParser) MustParse(inputFile string) *citymap.CityMap {

	cityMap, err := parser.TryParse(inputFile)
	utils.PanicIf(err)

	return cityMap
}

func (parser *MapParser) TryParse(inputFile string) (*citymap.CityMap, error) {

	lines := utils.MustReadLines(inputFile)

	if len(lines) == 0 {
		return nil, fmt.Errorf("empty input file: '%s'", inputFile)
	}

	width, height := len(lines[0]), len(lines)
	antennaMap := make(map[string][]antenna.Antenna)
	for lineIndex, line := range lines {
		if len(line) != width {
			return nil, fmt.Errorf(
				"can not parse map from file '%s': expect the length of line %d (%d) to be the same as the first line (%d)",
				inputFile, lineIndex+1, len(line), width,
			)
		}

		for colIndex, colRune := range line {
			if slices.Contains(parser.antennaLocationRunes, colRune) {
				antennaMap[string(colRune)] = append(
					antennaMap[string(colRune)],
					antenna.NewAntenna(
						string(colRune),
						location.NewLocation(lineIndex, colIndex),
					),
				)
			} else if slices.Contains(parser.emptyLocationRunes, colRune) {
				continue
			} else {
				return nil, fmt.Errorf(
					"can not parse map from file '%s': invalid rune '%c' at row %d, col %d",
					inputFile, colRune, lineIndex, colIndex,
				)
			}
		}
	}

	return &citymap.CityMap{
		Width:          width,
		Height:         height,
		AntennasByName: antennaMap,
	}, nil
}
