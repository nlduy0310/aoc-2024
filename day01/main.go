package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

const inputFilePath = "./input.txt"

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	leftArr, rightArr := make([]int, 0), make([]int, 0)
	file, err := os.Open(inputFilePath)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		var left, right int
		_, err = fmt.Sscanf(line, "%d %d", &left, &right)
		check(err)

		leftArr = append(leftArr, left)
		rightArr = append(rightArr, right)
	}

	err = scanner.Err()
	check(err)

	slices.Sort(leftArr)
	slices.Sort(rightArr)

	result1 := partOne(leftArr, rightArr)
	result2 := partTwo(leftArr, rightArr)
	fmt.Printf("Part one: %d\n", result1)
	fmt.Printf("Part two: %d\n", result2)
}

func partOne(leftArr, rightArr []int) int {
	total := 0
	for i := range len(leftArr) {
		diff := leftArr[i] - rightArr[i]
		if diff < 0 {
			diff = -diff
		}
		total += diff
	}
	return total
}

func partTwo(leftArr, rightArr []int) int {
	// fmt.Println(leftArr)
	// fmt.Println(rightArr)
	total := 0

	i, j, n := 0, 0, len(leftArr)
	for i < n && j < n {
		countLeft := 1
		for i+1 < n && leftArr[i+1] == leftArr[i] {
			i++
			countLeft++
		}
		// fmt.Printf("i moves to %d, countLeft = %d\n", i, countLeft)

		countRight := 0
		for j < n && rightArr[j] <= leftArr[i] {
			if rightArr[j] == leftArr[i] {
				countRight++
			}
			j++
		}
		// fmt.Printf("j moves to %d, countRight = %d\n", j, countRight)

		total += countLeft * leftArr[i] * countRight
		i++
	}

	return total
}
