package parttwo

import (
	"log"

	"github.com/nlduy0310/aoc-2024/day14/robot"
	"github.com/nlduy0310/aoc-2024/utils"
)

type Solver struct {
	inputFile     string
	robots        []*robot.Robot
	width, height int
}

func MustInitSolver(inputFile string, width, height int) Solver {

	lines := utils.MustReadLines(inputFile)

	robots := make([]*robot.Robot, len(lines))

	for idx := range len(lines) {
		robot, err := robot.TryParseFromString(lines[idx])
		if err != nil {
			log.Fatalf("can not parse robot from line %d: %s", idx, err.Error())
		}
		robots[idx] = robot
	}

	return Solver{
		inputFile: inputFile,
		robots:    robots,
		width:     width,
		height:    height,
	}
}

func (s Solver) Solve() int {

	seconds := 0

	for {
		for _, r := range s.robots {
			r.Move(robot.ClampXAfterMove(0, s.width-1), robot.ClampYAfterMove(0, s.height-1))
		}
		seconds++

		if s.checkChristmasTree() {
			break
		}
	}

	return seconds
}

func (s Solver) checkChristmasTree() bool {

	treeChecker := newTreeCheckerFromRobots(s.width, s.height, s.robots)

	return treeChecker.hasChristmasTree()
}
