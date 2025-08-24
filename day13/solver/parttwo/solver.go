package parttwo

import (
	"fmt"
	"log"
	"math"

	"github.com/nlduy0310/aoc-2024/day13/machine"
	"github.com/nlduy0310/aoc-2024/utils"
)

var linesPerMachine int = 3
var prizePointXOffset, prizePointYOffset int = 10000000000000, 10000000000000
var maximumButtonPress int = 100

type Solver struct {
	inputFile string
	machines  []*machine.Machine
}

func MustInitSolver(inputFile string) Solver {

	lines := utils.MustReadLines(inputFile)
	lines = utils.SliceFilter(lines, func(line string) bool { return len(line) > 0 })

	utils.Assert(len(lines) >= linesPerMachine, fmt.Sprintf("expected at least %d lines from input file '%s', got: %d", linesPerMachine, inputFile, len(lines)))

	machines := make([]*machine.Machine, len(lines)/linesPerMachine)
	for i := range len(lines) / linesPerMachine {
		machine, err := machine.TryParseFromLines([]string{lines[linesPerMachine*i], lines[linesPerMachine*i+1], lines[linesPerMachine*i+2]})
		if err != nil {
			log.Fatalf("can not parse machine from line %d to %d: %s", linesPerMachine*i, linesPerMachine*i+2, err.Error())
		}
		machine.PrizePoint.X += prizePointXOffset
		machine.PrizePoint.Y += prizePointYOffset
		machines[i] = machine
	}

	return Solver{
		inputFile: inputFile,
		machines:  machines,
	}
}

func (s Solver) Solve() int {

	ret := 0

	for _, machine := range s.machines {
		ret += s.solveMachine(machine)
	}

	return ret
}

func (s Solver) solveMachine(m *machine.Machine) int {

	utils.Assert(len(m.Buttons) == 2, fmt.Sprintf("expected 2 buttons per claw machine for this part, got: %d", len(m.Buttons)))
	firstButton, secondButton := m.Buttons[0], m.Buttons[1]

	crossProd := firstButton.XShift*secondButton.YShift - firstButton.YShift*secondButton.XShift

	if crossProd == 0 {
		return 0
	}

	firstButtonPressCount := float64(m.PrizePoint.X*secondButton.YShift-m.PrizePoint.Y*secondButton.XShift) / float64(crossProd)
	secondButtonPressCount := float64(firstButton.XShift*m.PrizePoint.Y-firstButton.YShift*m.PrizePoint.X) / float64(crossProd)

	if !(firstButtonPressCount == math.Trunc(firstButtonPressCount) && secondButtonPressCount == math.Trunc(secondButtonPressCount)) {
		return 0
	}

	// if int(firstButtonPressCount) > maximumButtonPress || int(secondButtonPressCount) > maximumButtonPress {
	// 	return 0
	// }

	return int(firstButtonPressCount)*firstButton.Cost + int(secondButtonPressCount)*secondButton.Cost
}
