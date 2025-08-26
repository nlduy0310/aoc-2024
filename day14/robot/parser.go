package robot

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/nlduy0310/aoc-2024/day14/position"
	"github.com/nlduy0310/aoc-2024/utils"
)

var robotPattern = regexp.MustCompile(`p=([-+]?\d+),([-+]?\d+) v=([-+]?\d+),([-+]?\d+)`)
var initialXGroupIndex, initialYGroupIndex = 1, 2
var velocityXGroupIndex, velocityYGroupIndex = 3, 4

func TryParseFromString(line string) (*Robot, error) {

	match := robotPattern.FindStringSubmatch(line)

	if match == nil {
		return nil, fmt.Errorf("can not parse robot from line '%s': not matching pattern", line)
	}

	initialX, err := strconv.Atoi(match[initialXGroupIndex])
	utils.PanicIf(err)
	initialY, err := strconv.Atoi(match[initialYGroupIndex])
	utils.PanicIf(err)
	velocityX, err := strconv.Atoi(match[velocityXGroupIndex])
	utils.PanicIf(err)
	velocityY, err := strconv.Atoi(match[velocityYGroupIndex])
	utils.PanicIf(err)

	initialPosition := position.NewPosition(initialX, initialY)
	ret := NewRobot(initialPosition, velocityX, velocityY)
	return &ret, nil
}
