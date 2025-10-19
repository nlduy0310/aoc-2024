package statecost

import (
	"fmt"

	"github.com/nlduy0310/aoc-2024/day16/state"
)

type StateCost struct {
	State state.State
	Cost  float64
}

func New(s state.State, cost float64) StateCost {
	return StateCost{
		State: s,
		Cost:  cost,
	}
}

func (sc StateCost) String() string {
	return fmt.Sprintf("StateCost[State=%s, Cost=%f]", sc.State, sc.Cost)
}
