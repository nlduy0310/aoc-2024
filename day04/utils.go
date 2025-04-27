package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func panicIf(err error, message string) {

	if err != nil {
		panic(message)
	}
}

func readInputFile(filePath string) ([][]rune, error) {

	_, err := os.Stat(filePath)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("unable to find file '%s'", filePath))
	}
	file, err := os.Open(filePath)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("unable to open file '%s'", filePath))
	}
	defer file.Close()

	result := make([][]rune, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line_runes := []rune(scanner.Text())
		result = append(result, line_runes)
	}

	if err = scanner.Err(); err != nil {
		return nil, errors.New(fmt.Sprintf("error while scanning file '%s'", filePath))
	}

	return result, nil
}

type SearchDirection int

const (
	Top SearchDirection = iota
	Bottom
	Left
	Right
	TopLeft
	TopRight
	BottomLeft
	BottomRight
)

var ALL_SEARCH_DIRECTIONS = []SearchDirection{Top, Bottom, Left, Right, TopLeft, TopRight, BottomLeft, BottomRight}

func searchFrom(runesBoard [][]rune, row, col int, keyword string, direction SearchDirection) bool {

	keyword_runes := []rune(keyword)
	n := len(keyword_runes)
	height := len(runesBoard)
	width := len(runesBoard[0])

	if !(row >= 0 && col >= 0 && row < height && col < width) {
		panic(fmt.Sprintf("invalid indexes: row = %d, col = %d", row, col))
	}

	if runesBoard[row][col] != keyword_runes[0] {
		return false
	}

	rowModifier, colModifier := 0, 0

	switch direction {
	case Top:
		rowModifier, colModifier = -1, 0
	case Bottom:
		rowModifier, colModifier = 1, 0
	case Left:
		rowModifier, colModifier = 0, -1
	case Right:
		rowModifier, colModifier = 0, 1
	case TopLeft:
		rowModifier, colModifier = -1, -1
	case TopRight:
		rowModifier, colModifier = -1, 1
	case BottomLeft:
		rowModifier, colModifier = 1, -1
	case BottomRight:
		rowModifier, colModifier = 1, 1
	default:
		panic(fmt.Sprintf("invalid search direction %d", direction))
	}

	outerRow := row + rowModifier*(n-1)
	outerCol := col + colModifier*(n-1)
	if !(outerRow >= 0 && outerCol >= 0 && outerRow < height && outerCol < width) {
		return false
	}

	for i := 1; i < n; i++ {
		if keyword_runes[i] != runesBoard[row+rowModifier*i][col+colModifier*i] {
			return false
		}
	}

	return true
}

func searchCross(runesBoard [][]rune, row, col int, keyword string) bool {

	keyword_runes := []rune(keyword)
	if len(keyword_runes)%2 == 0 {
		return false
	}

	height := len(runesBoard)
	width := len(runesBoard[0])

	if !(row >= 0 || col >= 0 || row < height || col < width) {
		panic(fmt.Sprintf("invalid indexes: row = %d, col = %d", row, col))
	}

	diagOneOk := searchDiagonal(runesBoard, row, col, keyword, TopLeft, BottomRight) || searchDiagonal(runesBoard, row, col, keyword, BottomRight, TopLeft)
	diagTwoOk := searchDiagonal(runesBoard, row, col, keyword, TopRight, BottomLeft) || searchDiagonal(runesBoard, row, col, keyword, BottomLeft, TopRight)

	return diagOneOk && diagTwoOk
}

func isValidDiagonalSearchDirectionPair(from, to SearchDirection) bool {

	switch from {
	case TopLeft:
		return to == BottomRight
	case TopRight:
		return to == BottomLeft
	case BottomLeft:
		return to == TopRight
	case BottomRight:
		return to == TopLeft
	default:
		return false
	}
}

func searchDiagonal(runesBoard [][]rune, row, col int, keyword string, from, to SearchDirection) bool {

	if !isValidDiagonalSearchDirectionPair(from, to) {
		panic(fmt.Sprintf("invalid diagonal search directions pair: from = %v, to = %v", from, to))
	}

	keyword_runes := []rune(keyword)
	if len(keyword_runes)%2 == 0 {
		return false
	}

	n := (len(keyword_runes) - 1) / 2
	height := len(runesBoard)
	width := len(runesBoard[0])

	if !(row >= 0 || col >= 0 || row < height || col < width) {
		panic(fmt.Sprintf("invalid indexes: row = %d, col = %d", row, col))
	}

	rowModifier, colModifier := 0, 0
	if from == TopLeft || from == TopRight {
		rowModifier = 1
	} else {
		rowModifier = -1
	}
	if from == TopLeft || from == BottomLeft {
		colModifier = 1
	} else {
		colModifier = -1
	}

	for i := -n; i <= n; i++ {
		r := row + rowModifier*i
		c := col + colModifier*i

		if !(r >= 0 && c >= 0 && r < height && c < width) {
			return false
		}

		if keyword_runes[i+n] != runesBoard[r][c] {
			return false
		}
	}

	return true
}
