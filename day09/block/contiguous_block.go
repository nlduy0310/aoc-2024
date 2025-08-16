package block

import (
	"fmt"

	"github.com/nlduy0310/aoc-2024/utils"
)

type ContiguousBlock struct {
	FileId int
	Size   int
}

func NewContiguousFileBlock(id int, size int) ContiguousBlock {

	utils.Assert(id >= 0, fmt.Sprintf("file id must not be negative, received %d", id))
	utils.Assert(size > 0, fmt.Sprintf("contiguous block creation with size %d is not permitted", size))

	return ContiguousBlock{
		FileId: id,
		Size:   size,
	}
}

func NewContiguousEmptyBlock(size int) ContiguousBlock {

	utils.Assert(size > 0, fmt.Sprintf("contiguous block creation with size %d is not permitted", size))

	return ContiguousBlock{
		FileId: -1,
		Size:   size,
	}
}

func (b ContiguousBlock) getFileId() int {

	return b.FileId
}
