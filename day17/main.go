package day17

import (
	"fmt"

	"github.com/nlduy0310/aoc-2024/day17/program"
)

func SolvePartOne(file string) {
	program, err := program.ParseFromFile(file)
	if err != nil {
		panic(fmt.Sprintf("can not parse program from file \"%s\": %s", file, err.Error()))
	}

	program.Execute()
	outp := program.GetOutput()

	fmt.Printf("Part one: %s\n", outp)
}
