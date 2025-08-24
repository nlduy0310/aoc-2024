package machine

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/nlduy0310/aoc-2024/day13/button"
	"github.com/nlduy0310/aoc-2024/day13/point"
	"github.com/nlduy0310/aoc-2024/utils"
)

var prizeLinePattern = regexp.MustCompile(`Prize:\s+X=([-+]?\d+), Y=([-+]?\d+)`)
var prizeXGroupIndex, prizeYGroupIndex = 1, 2

func validateShift(shiftAxis string, shiftValue int) error {

	if shiftValue <= 0 {
		return fmt.Errorf("expected %s shift value to be positive, got: %d", shiftAxis, shiftValue)
	}

	return nil
}

func parsePrizeLine(line string) (*point.Point, error) {

	match := prizeLinePattern.FindStringSubmatch(line)

	if match == nil {
		return nil, fmt.Errorf("can not parse prize from line '%s' for: not matching pattern", line)
	}

	prizeX, err := strconv.Atoi(match[prizeXGroupIndex])
	utils.PanicIf(err)
	prizeY, err := strconv.Atoi(match[prizeYGroupIndex])
	utils.PanicIf(err)

	ret := point.NewPoint(prizeX, prizeY)
	return &ret, nil
}

func TryParseFromLines(lines []string) (*Machine, error) {

	if len(lines) != 3 {
		return nil, fmt.Errorf("expected 3 lines for each machine, got: %d", len(lines))
	}

	button1, err := button.TryParseFromString(lines[0], 3)
	if err != nil {
		return nil, fmt.Errorf("can not parse button from first line: %s", err.Error())
	}

	button2, err := button.TryParseFromString(lines[1], 1)
	if err != nil {
		return nil, fmt.Errorf("can not parse button from second line: %s", err.Error())
	}

	for _, checkOpts := range []struct {
		axisName   string
		shiftValue int
	}{
		{"X", button1.XShift},
		{"Y", button1.YShift},
		{"X", button2.XShift},
		{"Y", button2.YShift},
	} {
		if err = validateShift(checkOpts.axisName, checkOpts.shiftValue); err != nil {
			return nil, err
		}
	}

	prizePoint, err := parsePrizeLine(lines[2])
	if err != nil {
		return nil, fmt.Errorf("can not parse prize from third line: %s", err.Error())
	}

	ret := NewMachine([]button.Button{*button1, *button2}, *prizePoint)
	return &ret, nil
}
