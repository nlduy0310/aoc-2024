package block

type Block interface {
	getFileId() int
}

func IsFileBlock(b Block) bool {

	return b.getFileId() >= 0
}

func IsEmptyBlock(b Block) bool {

	return b.getFileId() < 0
}
