package main

import "fmt"

// Supports simple types that can be copied easily
type TopologyNodeType interface {
	~int | ~float64 | ~string
}

type KahnsTopologySolver[T TopologyNodeType] struct {
	neighborsMap map[T][]T
}

func NewKahnsTopologySolver[T TopologyNodeType](neighborsMap map[T][]T) (*KahnsTopologySolver[T], error) {

	if len(neighborsMap) == 0 {
		return nil, fmt.Errorf("can not initialize solver with empty map")
	}

	copiedMap := make(map[T][]T)

	for key, value := range neighborsMap {
		copiedMap[key] = safeCopyList(value)
	}

	return &KahnsTopologySolver[T]{neighborsMap: copiedMap}, nil
}

func NewKahnsTopologySolverFromPairs[T TopologyNodeType](nodePairs [][2]T) (*KahnsTopologySolver[T], error) {

	if len(nodePairs) == 0 {
		return nil, fmt.Errorf("can not initialize solver with empty list of pairs")
	}

	neighborsMap := make(map[T][]T)

	for _, pair := range nodePairs {
		sourceNode, neighborNode := pair[0], pair[1]
		existingNeighbors, ok := neighborsMap[sourceNode]

		if ok {
			if !contains(existingNeighbors, neighborNode) {
				neighborsMap[sourceNode] = append(existingNeighbors, neighborNode)
			}
		} else {
			neighborsMap[sourceNode] = []T{neighborNode}
		}

		if _, ok := neighborsMap[neighborNode]; !ok {
			neighborsMap[neighborNode] = []T{}
		}
	}

	return &KahnsTopologySolver[T]{neighborsMap: neighborsMap}, nil
}

func (solver *KahnsTopologySolver[T]) Solve() [][]T {

	nodes := solver.nodes()
	inDegreesMap := make(map[T]int, len(nodes))

	for _, node := range nodes {
		inDegreesMap[node] = 0
	}

	for _, neighborNodes := range solver.neighborsMap {
		for _, neighborNode := range neighborNodes {
			inDegreesMap[neighborNode] = inDegreesMap[neighborNode] + 1
		}
	}

	ret := solver.solveRecursively([]T{}, inDegreesMap)

	return ret
}

func (solver *KahnsTopologySolver[T]) solveRecursively(current []T, inDegreesMap map[T]int) [][]T {

	if len(inDegreesMap) == 0 {
		return [][]T{safeCopyList(current)}
	}

	ret := make([][]T, 0)

	for node, inDegree := range inDegreesMap {
		if inDegree == 0 {

			newCurrent := append(safeCopyList(current), node)
			newInDegreesMap := safeCopyMap(inDegreesMap)
			delete(newInDegreesMap, node)

			for _, neighborNode := range solver.neighborsMap[node] {

				if _, ok := newInDegreesMap[neighborNode]; ok {
					newInDegreesMap[neighborNode] -= 1
				}
			}

			ret = append(ret, solver.solveRecursively(newCurrent, newInDegreesMap)...)
		}
	}

	return ret
}

func (solver *KahnsTopologySolver[T]) nodes() []T {

	ret := make([]T, 0, len(solver.neighborsMap))

	for node := range solver.neighborsMap {
		ret = append(ret, node)
	}

	return ret
}
