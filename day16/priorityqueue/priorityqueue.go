package priorityqueue

import (
	"fmt"
	"strings"

	"github.com/nlduy0310/aoc-2024/day16/statecost"
)

type PriorityQueue struct {
	heap []statecost.StateCost
}

func NewEmptyQueue() *PriorityQueue {
	return &PriorityQueue{
		heap: []statecost.StateCost{},
	}
}

func NewQueue(items []statecost.StateCost) *PriorityQueue {
	ret := PriorityQueue{
		heap: items,
	}
	heapify(ret.heap)
	return &ret
}

func (q *PriorityQueue) PrettyString() string {
	builder := strings.Builder{}

	builder.WriteString(fmt.Sprintf("PriorityQueue[len=%d]\n", len(q.heap)))
	builder.WriteString("--- HEAD ---\n")
	for _, item := range q.heap {
		builder.WriteString(item.String() + "\n")
	}
	builder.WriteString("--- TAIL ---")

	return builder.String()
}
