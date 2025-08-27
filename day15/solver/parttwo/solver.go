package parttwo

import (
	"fmt"

	"github.com/nlduy0310/aoc-2024/day15/direction"
	"github.com/nlduy0310/aoc-2024/day15/solver/parttwo/box"
	"github.com/nlduy0310/aoc-2024/day15/solver/parttwo/mapstate"
	"github.com/nlduy0310/aoc-2024/day15/solver/parttwo/parser"
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

func (s Solver) CalculateGPSSum() int {

	ret := 0

	for _, b := range s.mapState.Boxes {
		ret += 100*b.LeftPosition.Row + b.LeftPosition.Col
	}

	return ret
}

func (s Solver) moveBoxesRecursively(b *box.Box, d direction.Direction) ([]*box.Box, error) {

	affectedBoxes := make([]*box.Box, 0)
	for _, boundPos := range b.GetBound(d) {
		if !s.mapState.ContainsPosition(boundPos) {
			return nil, fmt.Errorf("can not move box %s", b.String())
		} else if s.mapState.IsWall(boundPos) {
			return nil, fmt.Errorf("can not move box %s", b.String())
		} else if affectedBox, err := s.mapState.TryGetBoxAt(boundPos); err == nil {
			affectedBoxes = append(affectedBoxes, affectedBox)
		}
	}

	ret := make([]*box.Box, 0)
	ret = append(ret, b)

	for _, affectedBox := range affectedBoxes {
		tmp, err := s.moveBoxesRecursively(affectedBox, d)
		if err != nil {
			return nil, fmt.Errorf("can not move box %s", b.String())
		}
		ret = append(ret, tmp...)
	}

	return ret, nil
}

func (s *Solver) moveInDirection(d direction.Direction) {

	nextPos := s.mapState.Robot.Moved(d, 1)

	if !s.mapState.ContainsPosition(nextPos) {
		return
	}

	if s.mapState.IsWall(nextPos) {
		return
	}

	if s.mapState.IsEmpty(nextPos) {
		s.mapState.Robot.Move(d, 1)
		return
	}

	nextBox, err := s.mapState.TryGetBoxAt(nextPos)
	utils.Assert(err == nil, fmt.Sprintf("next pos has to be a box: %s", nextPos.String()))
	boxesToMove, err := s.moveBoxesRecursively(nextBox, d)
	if err != nil {
		return
	}

	boxesToMove = utils.SliceUnique(boxesToMove)
	for _, boxToMove := range boxesToMove {
		boxToMove.Move(d, 1)
	}

	s.mapState.Robot.Move(d, 1)
}
