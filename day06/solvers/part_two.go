package solvers

import (
	"fmt"
	"slices"

	"github.com/nlduy0310/aoc-2024/day06/direction"
	"github.com/nlduy0310/aoc-2024/day06/guard"
	"github.com/nlduy0310/aoc-2024/day06/labmap"
	"github.com/nlduy0310/aoc-2024/day06/position"
	"github.com/nlduy0310/aoc-2024/utils"
)

type PartTwoSolver struct {
	inputFile         string
	lmap              labmap.Map
	lguard            guard.Guard
	obstaclePositions []position.Position
}

type VisitedObstacle struct {
	Position       position.Position
	VisitDirection direction.Direction
}

func findObstacles(lines []string) []position.Position {

	utils.Assert(lines != nil, "can not find obstacles from nil value")

	ret := make([]position.Position, 0)

	for rowIndex, line := range lines {
		for colIndex, cellRune := range line {
			cellValue := string(cellRune)
			if cellValue == "#" {
				ret = append(ret, position.NewPosition(rowIndex, colIndex))
			}
		}
	}

	return ret
}

func MustInitPartTwoSolver(inputFile string) *PartTwoSolver {

	lines := utils.MustReadLines(inputFile)

	lmap, err := labmap.NewMapFromLines(lines)
	utils.PanicIf(err)

	lguard := mustFindGuard(lines)

	return &PartTwoSolver{
		inputFile:         inputFile,
		lmap:              *lmap,
		lguard:            lguard,
		obstaclePositions: findObstacles(lines),
	}
}

func (solver *PartTwoSolver) isPositionInMap(position position.Position) bool {

	return utils.IsInRangeInclusive(position.Col, 0, solver.lmap.Cols()-1) &&
		utils.IsInRangeInclusive(position.Row, 0, solver.lmap.Rows()-1)
}

func (solver *PartTwoSolver) projectPosition(initialPosition position.Position, direction direction.Direction) position.Position {

	if solver.lmap.IsBlocked(initialPosition) {
		panic(fmt.Sprintf("can not project from a block position: %s", initialPosition))
	}

	currentPosition := initialPosition

	for solver.isPositionInMap(currentPosition) {
		nextPosition := currentPosition
		nextPosition.MoveInDirection(direction)

		if !solver.isPositionInMap(nextPosition) {
			currentPosition = nextPosition
		} else {
			if solver.lmap.IsBlocked(nextPosition) {
				break
			} else {
				currentPosition = nextPosition
			}
		}
	}

	return currentPosition
}

func (solver *PartTwoSolver) guardCanMoveForward() bool {

	nextPosition := solver.lguard.Position
	nextPosition.MoveInDirection(solver.lguard.Direction)

	if solver.isPositionInMap(nextPosition) {
		return !solver.lmap.IsBlocked(nextPosition)
	} else {
		return true
	}
}

func (solver *PartTwoSolver) turnGuardIfNecessary() (int, error) {

	attempts := 0
	for attempts <= 3 {
		if solver.guardCanMoveForward() {
			return attempts, nil
		} else {
			solver.lguard.Direction = solver.lguard.Direction.AfterTurn(direction.Right)
			attempts += 1
		}
	}

	return attempts, fmt.Errorf("guard can not move after 3 turn attempts")
}

// too lazy to write a fancy solution, so I just bruteforce
func (solver *PartTwoSolver) Solve() int {

	ret := make([]position.Position, 0)
	initialGuard := solver.lguard

	exitFound := false
	pathNumber := 1
	for !exitFound {
		pathNumber += 1
		_, err := solver.turnGuardIfNecessary()
		if err != nil {
			panic(fmt.Errorf("can not find a solution for part two, guard at %s, error: %s", solver.lguard.Position, err.Error()))
		}

		// get all the potential obstacle positions on the current path
		projectedPosition := solver.projectPosition(solver.lguard.Position, solver.lguard.Direction)
		potentialObstacles := make([]position.Position, 0)
		minRow, maxRow := utils.MinMax(solver.lguard.Position.Row, projectedPosition.Row)
		minCol, maxCol := utils.MinMax(solver.lguard.Position.Col, projectedPosition.Col)

		if solver.isPositionInMap(projectedPosition) {
			for row := minRow; row <= maxRow; row++ {
				for col := minCol; col <= maxCol; col++ {
					potentialObstacle := position.NewPosition(row, col)
					if potentialObstacle != solver.lguard.Position {
						potentialObstacles = append(potentialObstacles, potentialObstacle)
					}
				}
			}
		} else {
			for row := minRow; row <= maxRow; row++ {
				for col := minCol; col <= maxCol; col++ {
					potentialObstacle := position.NewPosition(row, col)
					if potentialObstacle != solver.lguard.Position && potentialObstacle != projectedPosition {
						potentialObstacles = append(potentialObstacles, potentialObstacle)
					}
				}
			}

			exitFound = true
		}

		// inspect potential obstacle to see if it causes loop
		for _, potentialObstacle := range potentialObstacles {
			mapCopy := solver.lmap.Copy()
			mapCopy.Block(potentialObstacle.Row, potentialObstacle.Col)

			caseSolver := PartOneSolver{
				inputFile: solver.inputFile,
				lmap:      *mapCopy,
				lguard:    initialGuard,
			}

			var causesLoop bool
			utils.SafeCall(
				func() {
					caseSolver.Solve()
					causesLoop = false
				},
				func(any) {
					causesLoop = true
				},
			)

			if causesLoop {
				if potentialObstacle != initialGuard.Position && !slices.Contains(ret, potentialObstacle) {
					ret = append(ret, potentialObstacle)
				}
			}
		}

		// next path
		if !exitFound {
			solver.lguard.Position = projectedPosition
		}
	}

	return len(ret)
}
