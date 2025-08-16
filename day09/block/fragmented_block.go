package block

import (
	"fmt"

	"github.com/nlduy0310/aoc-2024/utils"
)

type FragmentedBlock struct {
	FileId int
}

func NewFragmentedFileBlock(id int) FragmentedBlock {

	utils.Assert(id >= 0, fmt.Sprintf("file id can not be negative, received %d", id))

	return FragmentedBlock{id}
}

func NewFragmentedEmptyBlock() FragmentedBlock {

	return FragmentedBlock{-1}
}

func (b FragmentedBlock) getFileId() int {

	return b.FileId
}
