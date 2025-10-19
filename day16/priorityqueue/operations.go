package priorityqueue

import (
	"github.com/nlduy0310/aoc-2024/day16/statecost"
)

func heapify(items []statecost.StateCost) {
	for i := (len(items) - 1) / 2; i >= 0; i-- {
		heapifyDown(items, i)
	}
}

func heapifyDown(items []statecost.StateCost, idx int) {
	minItemIdx := idx
	leftIdx, rightIdx := 2*idx+1, 2*idx+2

	if leftIdx < len(items) && items[leftIdx].Cost < items[minItemIdx].Cost {
		minItemIdx = leftIdx
	}
	if rightIdx < len(items) && items[rightIdx].Cost < items[minItemIdx].Cost {
		minItemIdx = rightIdx
	}

	if minItemIdx != idx {
		items[idx], items[minItemIdx] = items[minItemIdx], items[idx]
		heapifyDown(items, minItemIdx)
	}
}

func heapifyUp(items []statecost.StateCost, idx int) {
	parentIdx := (idx - 1) / 2

	if parentIdx >= 0 && items[parentIdx].Cost > items[idx].Cost {
		items[idx], items[parentIdx] = items[parentIdx], items[idx]
		heapifyUp(items, parentIdx)
	}
}

func (p *PriorityQueue) Push(item statecost.StateCost) {
	p.heap = append(p.heap, item)
	heapifyUp(p.heap, len(p.heap)-1)
}

func (p *PriorityQueue) Pop() *statecost.StateCost {
	if len(p.heap) == 0 {
		return nil
	}

	ret := p.heap[0]
	p.heap[0], p.heap[len(p.heap)-1] = p.heap[len(p.heap)-1], p.heap[0]
	p.heap = p.heap[0 : len(p.heap)-1]
	heapifyDown(p.heap, 0)

	return &ret
}
