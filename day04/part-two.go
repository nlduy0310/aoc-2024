package main

func solvePartTwo(runesBoard [][]rune) int {

	result := 0
	for r := range len(runesBoard) {
		for c := range len(runesBoard[0]) {
			if ok := searchCross(runesBoard, r, c, "MAS"); ok {
				result++
			}
		}
	}

	return result
}
