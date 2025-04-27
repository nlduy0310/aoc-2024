package main

func solvePartOne(runesBoard [][]rune) int {

	result := 0
	for r := range len(runesBoard) {
		for c := range len(runesBoard[0]) {
			for _, direction := range ALL_SEARCH_DIRECTIONS {
				if ok := searchFrom(runesBoard, r, c, "XMAS", direction); ok {
					result++
				}
			}
		}
	}

	return result
}
