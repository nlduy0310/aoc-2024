package button

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/nlduy0310/aoc-2024/utils"
)

var pattern = regexp.MustCompile(`Button\s+(\w+):\s+X([-+]\d+),\s+Y([-+]\d+)`)
var nameGroupIndex, xShiftGroupIndex, yShiftGroupIndex int = 1, 2, 3

func TryParseFromString(str string, cost int) (*Button, error) {

	match := pattern.FindStringSubmatch(str)

	if match == nil {
		return nil, fmt.Errorf("the input string doesn't match the button pattern: '%s'", str)
	}

	name := match[nameGroupIndex]
	xShift, err := strconv.Atoi(match[xShiftGroupIndex])
	utils.PanicIf(err)
	yShift, err := strconv.Atoi(match[yShiftGroupIndex])
	utils.PanicIf(err)

	ret := NewButton(name, xShift, yShift, cost)
	return &ret, nil
}

func MustParseFromString(str string, cost int) *Button {

	buttonRef, err := TryParseFromString(str, cost)
	utils.PanicIf(err)

	return buttonRef
}
