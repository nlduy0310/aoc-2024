package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

const inputFilePath = "./input.txt"

type Violation int

const (
	NoViolation Violation = iota
	ViolateOrder
	ViolateAbs
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

type Report struct {
	levels []int
}

func NewReportFromString(line string) *Report {
	levels := strings.Split(line, " ")
	report := Report{
		levels: make([]int, len(levels)),
	}
	for i := range len(levels) {
		level := MustAsciiToInt(levels[i])
		report.levels[i] = level
	}
	return &report
}

func MustAsciiToInt(inp string) int {
	res, err := strconv.Atoi(inp)
	check(err)
	return res
}

func AbsInt(a, b int) int {
	res := a - b
	if res < 0 {
		res = -res
	}
	return res
}

func CheckPair(report Report, i, j int, diffMultiplier int) Violation {
	if i < 0 || i >= len(report.levels) || j < 0 || j >= len(report.levels) {
		return NoViolation
	}
	diff := report.levels[i] - report.levels[j]
	abs := AbsInt(report.levels[i], report.levels[j])
	// fmt.Println(report.levels)
	// fmt.Printf("Pair (%d, %d), diffMult = %d, diff = %d, abs = %d", report.levels[i], report.levels[j], diffMultiplier, diff, abs)
	if diff*diffMultiplier < 0 {
		// fmt.Println("\t---> ViolateOrder")
		return ViolateOrder
	}
	if abs < 1 || abs > 3 {
		// fmt.Println("\t---> ViolateAbs")
		return ViolateAbs
	}
	// fmt.Println("\t---> NoViolation")
	return NoViolation
}

func findDiffMultiplier(report Report) int {
	countIncreasing := 0
	for i := range len(report.levels) - 1 {
		if report.levels[i] < report.levels[i+1] {
			countIncreasing++
		}
		if countIncreasing >= len(report.levels)/2 {
			return -1
		}
	}
	return 1
}

func main() {
	f, err := os.Open(inputFilePath)
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	reports := make([]Report, 0)

	for scanner.Scan() {
		line := scanner.Text()
		reports = append(reports, *NewReportFromString(line))
	}

	err = scanner.Err()
	check(err)

	resultOne := partOne(reports)
	resultTwo := partTwo(reports)
	fmt.Printf("Part one: %d\n", resultOne)
	fmt.Printf("Part two: %d\n", resultTwo)
}

func partOne(reports []Report) int {
	count := len(reports)
	for _, report := range reports {
		lastDiff := 0
		for i := range len(report.levels) - 1 {
			diff := report.levels[i] - report.levels[i+1]
			abs := AbsInt(report.levels[i], report.levels[i+1])
			if diff*lastDiff < 0 || abs < 1 || abs > 3 {
				count--
				break
			}
			lastDiff = diff
		}
	}
	return count
}

func partTwo(reports []Report) int {
	count := len(reports)
	for _, report := range reports {
		flag := false
		diffMultiplier := findDiffMultiplier(report)
		for i := 0; i < len(report.levels)-1; i++ {
			violation := CheckPair(report, i, i+1, diffMultiplier)
			if violation == NoViolation {
				continue
			} else {
				if flag {
					count--
					break
				}
				v1 := CheckPair(report, i-1, i+1, diffMultiplier)
				v2 := CheckPair(report, i, i+2, diffMultiplier)
				if v1 != NoViolation && v2 == NoViolation {
					report.levels = slices.Delete(report.levels, i+1, i+2)
					i -= 1
				} else if v1 == NoViolation && v2 != NoViolation {
					report.levels = slices.Delete(report.levels, i, i+1)
					i -= 2
				} else if v1 == NoViolation && v2 == NoViolation {
					if CheckPair(report, i+1, i+2, diffMultiplier) == NoViolation {
						report.levels = slices.Delete(report.levels, i, i+1)
						i -= 2
					} else if CheckPair(report, i-1, i, diffMultiplier) == NoViolation {
						report.levels = slices.Delete(report.levels, i+1, i+2)
						i -= 1
					}
				} else {
					count--
					break
				}
				flag = true
			}
		}
	}
	return count
}
