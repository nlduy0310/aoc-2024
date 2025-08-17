package topographicmap

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/nlduy0310/aoc-2024/utils"
	runes_utils "github.com/nlduy0310/aoc-2024/utils/runes"
)

type TopographicMap struct {
	cells  [][]int
	width  int
	height int
}

func MustParseFromFile(inputFile string) TopographicMap {

	lines := utils.MustReadLines(inputFile)
	utils.Assert(len(lines) > 0, fmt.Sprintf("input file is empty '%s'", inputFile))

	width := len(lines[0])
	height := len(lines)

	cells := make([][]int, 0, height)

	for idx, line := range lines {
		utils.Assert(len(line) == width, fmt.Sprintf("error parsing input file '%s': line %d has inconsistent length with first line", inputFile, idx+1))

		lineCells := make([]int, 0, width)
		for _, digitRune := range line {
			utils.Assert(slices.Contains(runes_utils.Digits, digitRune), fmt.Sprintf("error parsing input file '%s': found invalid character '%c'", inputFile, digitRune))

			digit, _ := strconv.Atoi(string(digitRune))
			lineCells = append(lineCells, digit)
		}
		cells = append(cells, lineCells)
	}

	return TopographicMap{
		cells:  cells,
		width:  width,
		height: height,
	}
}

func (m *TopographicMap) PrettyString() string {

	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("TopographicMap: width=%d, height=%d\n", m.width, m.height))

	for _, row := range m.cells {
		for _, cellValue := range row {
			builder.WriteString(strconv.Itoa(cellValue))
		}
		builder.WriteString("\n")
	}

	return builder.String()
}
