package parser

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nlduy0310/aoc-2024/day11/stone"
	"github.com/nlduy0310/aoc-2024/utils"
)

func TryParseFromLine(line string) ([]stone.Stone, error) {

	tokens := strings.Split(line, " ")
	ret := make([]stone.Stone, 0, len(tokens))

	for _, token := range tokens {
		val, err := strconv.Atoi(token)
		if err != nil {
			return nil, fmt.Errorf("can not convert '%s' to int: %s", token, err.Error())
		}

		ret = append(ret, stone.Stone{Val: val})
	}

	return ret, nil
}

func MustParseFromFile(file string) []stone.Stone {

	lines := utils.MustReadLines(file)
	utils.Assert(len(lines) > 0, fmt.Sprintf("empty input file '%s'", file))

	return MustParseFromLine(lines[0])
}

func MustParseFromLine(line string) []stone.Stone {

	stones, err := TryParseFromLine(line)
	utils.PanicIf(err)

	return stones
}
