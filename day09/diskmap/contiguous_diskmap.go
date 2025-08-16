package diskmap

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/nlduy0310/aoc-2024/day09/block"
	"github.com/nlduy0310/aoc-2024/utils"
	runes_utils "github.com/nlduy0310/aoc-2024/utils/runes"
)

type ContiguousDiskMap struct {
	Blocks []block.ContiguousBlock
}

func MustParseContiguousFromLine(line string) ContiguousDiskMap {

	line = strings.TrimSpace(line)

	blocks := make([]block.ContiguousBlock, 0, len(line))

	for idx, digitRune := range line {
		if !slices.Contains(runes_utils.Digits, digitRune) {
			panic(fmt.Sprintf("invalid digit rune '%c'", digitRune))
		}

		isFileSizeDigit := idx%2 == 0
		size, _ := strconv.Atoi(string(digitRune))
		if size == 0 {
			continue
		}

		if isFileSizeDigit {
			fileIndex := idx / 2
			blocks = append(blocks, block.NewContiguousFileBlock(fileIndex, size))
		} else {
			blocks = append(blocks, block.NewContiguousEmptyBlock(size))
		}
	}

	return ContiguousDiskMap{
		Blocks: blocks,
	}
}

func (m *ContiguousDiskMap) concatenateEmptyContiguousBlocks() {

	leftCursor := 0
	for {
		for leftCursor < len(m.Blocks) && !block.IsEmptyBlock(m.Blocks[leftCursor]) {
			leftCursor++
		}
		if leftCursor >= len(m.Blocks) {
			break
		}

		rightCursor := leftCursor + 1
		for ; rightCursor < len(m.Blocks) && block.IsEmptyBlock(m.Blocks[rightCursor]); rightCursor++ {
			continue
		}
		if rightCursor >= len(m.Blocks) {
			break
		}

		if rightCursor-leftCursor == 1 {
			leftCursor = rightCursor
			continue
		}

		totalEmptyBlockSize := 0
		for i := leftCursor; i < rightCursor; i++ {
			totalEmptyBlockSize += m.Blocks[i].Size
		}

		m.Blocks = slices.Replace(m.Blocks, leftCursor, rightCursor, block.NewContiguousEmptyBlock(totalEmptyBlockSize))
		leftCursor += 1
	}
}

func (m *ContiguousDiskMap) Compact() {

	largestFileIdIndex := utils.SliceFindLast(
		m.Blocks,
		func(b block.ContiguousBlock) bool { return block.IsFileBlock(b) },
	)
	utils.Assert(largestFileIdIndex >= 0, "can not find any file in diskmap")
	largestFileId := m.Blocks[largestFileIdIndex].FileId

	for fileId := largestFileId; fileId >= 0; fileId-- {

		fileContiguousBlockIndex := utils.SliceFindFirst(
			m.Blocks,
			func(b block.ContiguousBlock) bool { return b.FileId == fileId },
		)
		utils.Assert(fileContiguousBlockIndex >= 0, fmt.Sprintf("can not find block with file id %d", fileId))

		fileBlock := m.Blocks[fileContiguousBlockIndex]
		for i := 0; i < fileContiguousBlockIndex; i++ {
			if !block.IsEmptyBlock(m.Blocks[i]) {
				continue
			}
			if !(m.Blocks[i].Size >= fileBlock.Size) {
				continue
			}

			m.Blocks[fileContiguousBlockIndex] = block.NewContiguousEmptyBlock(fileBlock.Size)
			if m.Blocks[i].Size == fileBlock.Size {
				m.Blocks = slices.Replace(m.Blocks, i, i+1, fileBlock)
			} else {
				sizeDiff := m.Blocks[i].Size - fileBlock.Size
				m.Blocks = slices.Replace(m.Blocks, i, i+1, fileBlock, block.NewContiguousEmptyBlock(sizeDiff))
			}
			m.concatenateEmptyContiguousBlocks()
			break
		}
	}
}

func (m *ContiguousDiskMap) PrettyString() string {

	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("ContiguousDiskMap: len=%d\n", len(m.Blocks)))

	for _, b := range m.Blocks {
		for range b.Size {
			if block.IsEmptyBlock(b) {
				builder.WriteString(".")
			} else {
				builder.WriteString(strconv.Itoa(b.FileId))
			}
		}
		builder.WriteString(" ")
	}

	builder.WriteString("\n")

	return builder.String()
}

func (m *ContiguousDiskMap) CalculateChecksum() int {

	ret := 0

	fragmentedIndex := 0
	for _, b := range m.Blocks {

		if !block.IsEmptyBlock(b) {
			for i := 0; i < b.Size; i++ {
				ret += b.FileId * (fragmentedIndex + i)
			}
		}
		fragmentedIndex += b.Size
	}

	return ret
}
