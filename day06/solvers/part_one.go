package solvers

import (
	"fmt"

	"github.com/nlduy0310/aoc-2024/day06/direction"
	"github.com/nlduy0310/aoc-2024/day06/guard"
	"github.com/nlduy0310/aoc-2024/day06/labmap"
	"github.com/nlduy0310/aoc-2024/day06/position"
	"github.com/nlduy0310/aoc-2024/utils"
)

type PartOneSolver struct {
	inputFile string
	lmap      labmap.Map
	lguard    guard.Guard
}

func MustInitPartOneSolver(inputFile string) *PartOneSolver {

	lines := utils.MustReadLines(inputFile)

	lmap, err := labmap.NewMapFromLines(lines)
	utils.PanicIf(err)

	lguard := mustFindGuard(lines)

	return &PartOneSolver{
		inputFile: inputFile,
		lmap:      *lmap,
		lguard:    lguard,
	}
}

func (solver *PartOneSolver) guardIsInMap() bool {

	return solver.isPositionInMap(solver.lguard.Position)
}

func (solver *PartOneSolver) isPositionInMap(position position.Position) bool {

	return utils.IsInRangeInclusive(position.Row, 0, solver.lmap.Rows()-1) &&
		utils.IsInRangeInclusive(position.Col, 0, solver.lmap.Cols()-1)
}

func (solver *PartOneSolver) guardCanMoveForward() bool {

	nextGuardPosition := solver.lguard.Position
	nextGuardPosition.MoveInDirection(solver.lguard.Direction)

	if solver.isPositionInMap(nextGuardPosition) {
		return !solver.lmap.IsBlocked(nextGuardPosition)
	} else {
		return true
	}
}

func (solver *PartOneSolver) projectPosition(initialPosition position.Position, direction direction.Direction) position.Position {

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

func (solver *PartOneSolver) tryTurningRight() (int, error) {

	attempts := 0
	for attempts <= 3 {
		if solver.guardCanMoveForward() {
			break
		}

		solver.lguard.Direction = solver.lguard.Direction.AfterTurn(direction.Right)
		attempts += 1
	}

	if attempts > 3 {
		return attempts, fmt.Errorf("can not turn guard after 3 attempts")
	}

	return attempts, nil
}

func (solver *PartOneSolver) Solve() int {

	visitedObstaclesMap := make(map[VisitedObstacle]bool)
	visitedCells := make(map[position.Position]bool)
	visitedCells[solver.lguard.Position] = true

	exitFound := false
	iteration := 1
	for !exitFound {

		iteration += 1

		_, err := solver.tryTurningRight()
		utils.PanicIf(err)

		projectedPosition := solver.projectPosition(solver.lguard.Position, solver.lguard.Direction)

		if solver.isPositionInMap(projectedPosition) {
			obstaclePosition := projectedPosition
			obstaclePosition.MoveInDirection(solver.lguard.Direction)
			if _, visited := visitedObstaclesMap[VisitedObstacle{Position: obstaclePosition, VisitDirection: solver.lguard.Direction}]; visited {
				// loop detected
				panic(fmt.Errorf("loop detected at: %s", obstaclePosition))
			} else {
				visitedObstaclesMap[VisitedObstacle{Position: obstaclePosition, VisitDirection: solver.lguard.Direction}] = true
			}

		} else {
			projectedPosition.MoveInDirection(solver.lguard.Direction.Opposite())
			exitFound = true
		}

		minRow, maxRow := utils.MinMax(solver.lguard.Position.Row, projectedPosition.Row)
		minCol, maxCol := utils.MinMax(solver.lguard.Position.Col, projectedPosition.Col)

		for row := minRow; row <= maxRow; row++ {
			for col := minCol; col <= maxCol; col++ {
				visitedCells[position.NewPosition(row, col)] = true
			}
		}

		if !exitFound {
			solver.lguard.Position = projectedPosition
		}
	}

	return len(visitedCells)
}
