package partone

import (
	"github.com/nlduy0310/aoc-2024/day15/direction"
	"github.com/nlduy0310/aoc-2024/day15/position"
	"github.com/nlduy0310/aoc-2024/day15/solver/partone/mapstate"
	"github.com/nlduy0310/aoc-2024/day15/solver/partone/parser"
	"github.com/nlduy0310/aoc-2024/utils"
)

type Solver struct {
	inputFile  string
	mapState   *mapstate.MapState
	directions []direction.Direction
}

func MustInitSolver(inputFile string) Solver {

	mapState, directions, err := parser.TryParseFromFile(inputFile)
	utils.ExitIf(err)

	return Solver{
		inputFile:  inputFile,
		mapState:   mapState,
		directions: directions,
	}
}

func (s Solver) Solve() int {

	for _, d := range s.directions {
		s.moveInDirection(d)
	}

	return s.CalculateGPSSum()
}

func (s *Solver) moveInDirection(d direction.Direction) {

	boxesToMove := make([]*position.Position, 0)

	currentPosition := s.mapState.Robot
	for {
		currentPosition = currentPosition.Moved(d, 1)

		if !s.mapState.ContainsPosition(currentPosition) {
			break
		} else if s.mapState.IsWall(currentPosition) {
			break
		} else if box := s.mapState.TryGetBox(currentPosition); box != nil {
			boxesToMove = append(boxesToMove, box)
		} else {
			if len(boxesToMove) == 0 {
				s.mapState.Robot.Move(d, 1)
			}
			break
		}
	}

	if len(boxesToMove) == 0 {
		return
	}

	if !s.mapState.ContainsPosition(currentPosition) || s.mapState.IsWall(currentPosition) {
		return
	}

	utils.Assert(!s.mapState.IsWall(currentPosition), "if not out of map or wall encountered, then the final position must be empty")
	for _, boxToMove := range boxesToMove {
		boxToMove.Move(d, 1)
	}
	s.mapState.Robot.Move(d, 1)
}

func (s Solver) CalculateGPSSum() int {

	ret := 0

	for _, box := range s.mapState.Boxes {
		ret += 100*box.Row + box.Col
	}

	return ret
}
