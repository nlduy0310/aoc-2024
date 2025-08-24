package machine

import (
	"fmt"
	"strings"

	"github.com/nlduy0310/aoc-2024/day13/button"
	"github.com/nlduy0310/aoc-2024/day13/point"
	"github.com/nlduy0310/aoc-2024/utils"
)

type Machine struct {
	Buttons    []button.Button
	PrizePoint point.Point
}

func NewMachine(buttons []button.Button, prizePoint point.Point) Machine {

	return Machine{
		Buttons:    buttons,
		PrizePoint: prizePoint,
	}
}

func (m *Machine) String() string {

	builder := strings.Builder{}
	builder.WriteString("Machine[")

	buttonStrings := utils.SliceMap(m.Buttons, func(b button.Button) string { return b.String() })
	builder.WriteString(fmt.Sprintf("buttons=[%s], ", strings.Join(buttonStrings, ", ")))

	builder.WriteString(fmt.Sprintf("prizePoint=%s", m.PrizePoint.String()))

	builder.WriteString("]")

	return builder.String()
}
