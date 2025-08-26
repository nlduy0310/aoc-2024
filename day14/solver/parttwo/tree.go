package parttwo

import (
	"github.com/nlduy0310/aoc-2024/day14/position"
	"github.com/nlduy0310/aoc-2024/day14/robot"
	"github.com/nlduy0310/aoc-2024/utils"
)

type christmasTreeChecker struct {
	positions           map[position.Position]struct{}
	mapWidth, mapHeight int
}

func (c *christmasTreeChecker) addPosition(p position.Position) {

	c.positions[p] = struct{}{}
}

func (c *christmasTreeChecker) hasPosition(p position.Position) bool {

	_, ok := c.positions[p]

	return ok
}

func (c *christmasTreeChecker) hasValidPosition(p position.Position) bool {

	return c.hasPosition(p) &&
		utils.IsInRangeInclusive(p.X, 0, c.mapWidth-1) &&
		utils.IsInRangeInclusive(p.Y, 0, c.mapHeight-1)
}

func newTreeChecker(width, height int) christmasTreeChecker {

	return christmasTreeChecker{
		positions: make(map[position.Position]struct{}),
		mapWidth:  width,
		mapHeight: height,
	}
}

func newTreeCheckerFromRobots(width, height int, robots []*robot.Robot) christmasTreeChecker {

	ret := newTreeChecker(width, height)

	for _, r := range robots {
		ret.addPosition(r.CurrentPosition)
	}

	return ret
}

func (c christmasTreeChecker) hasChristmasTree() bool {

	for potentialRoot := range c.positions {
		if c.checkFromPotentialRoot(potentialRoot) {
			return true
		}
	}

	return false
}

func (c christmasTreeChecker) checkFromPotentialRoot(root position.Position) bool {

	// check immediate upper
	upperCenter := root.GetUp(1)
	if !c.hasValidPosition(upperCenter) {
		return false
	}

	i := 1
	for ; c.hasValidPosition(upperCenter.GetLeft(i)) && c.hasValidPosition(upperCenter.GetRight(i)); i++ {
		continue
	}

	if i <= 2 {
		return false
	}

	base := upperCenter
	for w := 0; w < i-1; w++ {
		base = base.GetUp(1)
		if !c.hasValidPosition(base) {
			return false
		}

		for x := 0; x < w; x++ {
			if !(c.hasValidPosition(base.GetLeft(x)) && c.hasValidPosition(base.GetRight(x))) {
				return false
			}
		}
	}

	// println("found christmas tree from root: ", root.String())
	// c.visualize()
	return true
}

// func (c christmasTreeChecker) visualize() {

// 	builder := strings.Builder{}

// 	builder.WriteString("--- VISUALIZER ---\n")

// 	for y := 0; y < c.mapHeight; y++ {
// 		for x := 0; x < c.mapWidth; x++ {
// 			if c.hasPosition(position.NewPosition(x, y)) {
// 				builder.WriteRune('X')
// 			} else {
// 				builder.WriteRune('.')
// 			}
// 		}
// 		builder.WriteRune('\n')
// 	}

// 	print(builder.String())
// }
