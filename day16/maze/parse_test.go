package maze_test

import (
	"testing"

	"github.com/nlduy0310/aoc-2024/day16/maze"
)

const (
	mazeFile = "./testdata/maze"
)

func TestParseFromFile(t *testing.T) {
	_, err := maze.ParseFromFile(mazeFile)
	if err != nil {
		t.Errorf("test parsing file \"%s\" failed: %s", mazeFile, err.Error())
	}
}
