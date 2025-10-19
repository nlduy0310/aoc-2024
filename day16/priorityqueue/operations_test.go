package priorityqueue

import (
	"testing"

	"github.com/nlduy0310/aoc-2024/day16/direction"
	"github.com/nlduy0310/aoc-2024/day16/position"
	"github.com/nlduy0310/aoc-2024/day16/state"
	"github.com/nlduy0310/aoc-2024/day16/statecost"
)

func TestOperations(t *testing.T) {
	pq := NewEmptyQueue()

	p := position.NewPosition(0, 0)
	d := direction.North

	pq.Push(statecost.New(state.NewState(p, d.TurnedLeft()), 1000))
	pq.Push(statecost.New(state.NewState(p, d), 1))

	println(pq.PrettyString())

	var top *statecost.StateCost = pq.Pop()
	for ; top != nil; top = pq.Pop() {
		println(top.String())
	}
}
