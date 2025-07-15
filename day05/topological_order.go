package main

import "fmt"

type TopologicalOrder struct {
	order []int
}

func (o *TopologicalOrder) String() string {
	return fmt.Sprintf("TopologicalOrder[order=%v]", o.order)
}

func NewTopologicalOrder(order []int) (*TopologicalOrder, error) {

	if order == nil {
		return nil, fmt.Errorf("can not initialize a topological order with nil")
	}

	if len(order) == 0 {
		return nil, fmt.Errorf("can not initialize a topological order with an empty list")
	}

	return &TopologicalOrder{order: safeCopyList(order)}, nil
}

func (order *TopologicalOrder) Matches(targetOrder []int) bool {

	finalTargetOrder := reserve(targetOrder, order.order)
	finalOrder := reserve(order.order, finalTargetOrder)

	pointer1, pointer2 := 0, 0

	// handle the case where elements of targetOrder are not unique
	for pointer2 < len(finalTargetOrder) {
		for pointer1 < len(finalOrder) {
			if finalOrder[pointer1] == finalTargetOrder[pointer2] {
				break
			}
			pointer1 += 1
		}

		if pointer1 >= len(finalOrder) {
			return false
		}

		pointer2 += 1
	}

	return true
}

func (order *TopologicalOrder) hasWildcard(targetOrder []int) bool {

	for _, val := range targetOrder {
		if !contains(order.order, val) {
			return true
		}
	}

	return false
}

func (order *TopologicalOrder) Fix(targetOrder []int) ([]int, error) {

	if order.hasWildcard(targetOrder) {
		return nil, fmt.Errorf("can not determine a fix for %v, at least one element is unbound", targetOrder)
	}

	finalOrder := reserve(order.order, targetOrder)
	// len(finalOrder) <= len(finalTargetOrder)

	ret := make([]int, 0, len(targetOrder))
	for _, val := range finalOrder {
		valCount := count(targetOrder, val) // >= 1
		for range valCount {
			ret = append(ret, val)
		}
	}

	return ret, nil
}
