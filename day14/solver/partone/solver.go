package partone

import (
	"log"

	"github.com/nlduy0310/aoc-2024/day14/robot"
	"github.com/nlduy0310/aoc-2024/utils"
)

type Solver struct {
	inputFile     string
	robots        []*robot.Robot
	width, height int
	seconds       int
}

func MustInitSolver(inputFile string, width, height int, seconds int) Solver {

	lines := utils.MustReadLines(inputFile)

	robots := make([]*robot.Robot, len(lines))

	for i := range len(lines) {
		robot, err := robot.TryParseFromString(lines[i])
		if err != nil {
			log.Fatalf("can not parse robot from line %d: %s", i, err.Error())
		}
		robots[i] = robot
	}

	return Solver{
		inputFile: inputFile,
		robots:    robots,
		width:     width,
		height:    height,
		seconds:   seconds,
	}
}

func (s Solver) countQuadrants() []int {

	type Quadrant struct {
		xMin, xMax, yMin, yMax int
	}

	contains := func(q Quadrant, r robot.Robot) bool {

		return utils.IsInRangeInclusive(r.CurrentPosition.X, q.xMin, q.xMax) &&
			utils.IsInRangeInclusive(r.CurrentPosition.Y, q.yMin, q.yMax)
	}

	quadrantXLength, quadrantYLength := s.width/2, s.height/2

	quadrants := []Quadrant{
		{0, 0 + (quadrantXLength - 1), 0, 0 + (quadrantYLength - 1)},
		{0, 0 + (quadrantXLength - 1), s.height - quadrantYLength, s.height - 1},
		{s.width - quadrantXLength, s.width - 1, s.height - quadrantYLength, s.height - 1},
		{s.width - quadrantXLength, s.width - 1, 0, 0 + (quadrantYLength - 1)},
	}

	ret := utils.SliceInit(4, 0)

	for _, r := range s.robots {
		for idx, quadrant := range quadrants {
			if contains(quadrant, *r) {
				ret[idx] += 1
				break
			}
		}
	}

	return ret
}

func (s Solver) Solve() int {

	for _, r := range s.robots {
		for range s.seconds {
			r.Move(robot.ClampXAfterMove(0, s.width-1), robot.ClampYAfterMove(0, s.height-1))
		}
	}

	quadrantCounts := s.countQuadrants()
	safetyFactor := 1

	for _, count := range quadrantCounts {
		safetyFactor *= count
	}

	return safetyFactor
}
