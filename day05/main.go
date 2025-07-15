package main

import "fmt"

var PART_1_INPUT string = "./day05/data/input-1"
var PART_2_INPUT string = "./day05/data/input-2"

func partOne() {

	parser, err := NewParser(PART_1_INPUT)
	panicIf(err)

	rulePairs, updates, err := parser.Parse()
	panicIf(err)

	res := 0
	for _, update := range updates {
		relatedRulePairs := filter(rulePairs, func(p RulePair) bool {
			return contains(update.Pages, p.Before) && contains(update.Pages, p.After)
		})

		topologySolver, err := NewKahnsTopologySolverFromPairs(mapTo(relatedRulePairs, func(p RulePair) [2]int {
			return [2]int{p.Before, p.After}
		}))
		panicIf(err)

		topologicalOrdersAsIntSlices := topologySolver.Solve()
		topologicalOrders := mapTo(topologicalOrdersAsIntSlices, func(slice []int) *TopologicalOrder {
			ret, err := NewTopologicalOrder(slice)
			panicIf(err)
			return ret
		})

		for _, topologicalOrder := range topologicalOrders {
			if topologicalOrder.Matches(update.Pages) {
				res += update.Pages[len(update.Pages)/2]
			}
		}
	}

	fmt.Printf("Part one's result: %d\n", res)
}

func partTwo() {

	parser, err := NewParser(PART_2_INPUT)
	panicIf(err)

	rulePairs, updates, err := parser.Parse()
	panicIf(err)

	res := 0
	for _, update := range updates {
		relatedRulePairs := filter(rulePairs, func(p RulePair) bool {
			return contains(update.Pages, p.Before) && contains(update.Pages, p.After)
		})

		topologySolver, err := NewKahnsTopologySolverFromPairs(mapTo(relatedRulePairs, func(p RulePair) [2]int {
			return [2]int{p.Before, p.After}
		}))
		panicIf(err)

		topologicalOrdersAsIntSlices := topologySolver.Solve()
		topologicalOrders := mapTo(topologicalOrdersAsIntSlices, func(slice []int) *TopologicalOrder {
			ret, err := NewTopologicalOrder(slice)
			panicIf(err)
			return ret
		})

		for _, topologicalOrder := range topologicalOrders {
			if !topologicalOrder.Matches(update.Pages) {
				fixedUpdate, err := topologicalOrder.Fix(update.Pages)
				panicIf(err)
				res += fixedUpdate[len(fixedUpdate)/2]
			}
		}
	}

	fmt.Printf("Part two's result: %d\n", res)
}

func main() {

	partOne()
	partTwo()
}
