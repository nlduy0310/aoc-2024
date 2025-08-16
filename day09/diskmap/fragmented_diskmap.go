package diskmap

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/nlduy0310/aoc-2024/day09/block"
	runes_utils "github.com/nlduy0310/aoc-2024/utils/runes"
)

type FragmentedDiskMap struct {
	Blocks []block.FragmentedBlock
}

func MustParseFragmentedFromLine(line string) FragmentedDiskMap {

	line = strings.TrimSpace(line)

	counts := make([]int, 0, len(line))
	sum := 0

	for _, digitRune := range line {
		if slices.Contains(runes_utils.Digits, digitRune) {
			digit, _ := strconv.Atoi(string(digitRune))
			counts = append(counts, digit)
			sum += digit
		} else {
			panic(fmt.Sprintf("invalid digit found: '%c'", digitRune))
		}
	}

	blocks := make([]block.FragmentedBlock, 0, sum)

	for idx, count := range counts {
		var blockTemplate block.FragmentedBlock
		if idx%2 == 0 {
			blockTemplate = block.NewFragmentedFileBlock(idx / 2)
		} else {
			blockTemplate = block.NewFragmentedEmptyBlock()
		}
		for i := 0; i < count; i++ {
			blocks = append(blocks, blockTemplate)
		}
	}

	return FragmentedDiskMap{
		Blocks: blocks,
	}
}

func (m *FragmentedDiskMap) PrettyString() string {

	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("FragmentedDiskMap: len=%d\n", len(m.Blocks)))

	for _, b := range m.Blocks {
		if block.IsEmptyBlock(b) {
			builder.WriteString(".")
		} else {
			builder.WriteString(strconv.Itoa(b.FileId))
		}
		builder.WriteString(" ")
	}

	builder.WriteString("\n")

	return builder.String()
}

func (m *FragmentedDiskMap) Compact() {

	var leftmostEmptyBlockIndex, rightmostFileBlockIndex int = 0, len(m.Blocks) - 1

	for {
		for ; !block.IsEmptyBlock(m.Blocks[leftmostEmptyBlockIndex]) && leftmostEmptyBlockIndex < len(m.Blocks); leftmostEmptyBlockIndex++ {
			continue
		}

		for ; !block.IsFileBlock(m.Blocks[rightmostFileBlockIndex]) && rightmostFileBlockIndex >= 0; rightmostFileBlockIndex-- {
			continue
		}

		if leftmostEmptyBlockIndex > rightmostFileBlockIndex || leftmostEmptyBlockIndex >= len(m.Blocks) || rightmostFileBlockIndex < 0 {
			break
		}

		m.Blocks[leftmostEmptyBlockIndex], m.Blocks[rightmostFileBlockIndex] = m.Blocks[rightmostFileBlockIndex], m.Blocks[leftmostEmptyBlockIndex]
	}
}

func (m *FragmentedDiskMap) CalculateChecksum() int {

	ret := 0

	for idx, b := range m.Blocks {
		if block.IsEmptyBlock(b) {
			continue
		}

		ret += idx * b.FileId
	}

	return ret
}
