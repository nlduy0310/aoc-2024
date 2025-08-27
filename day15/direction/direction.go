package direction

import (
	"fmt"
	"log"
)

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

var stringMap = map[Direction]string{
	Up:    "Direction[Up]",
	Down:  "Direction[Down]",
	Left:  "Direction[Left]",
	Right: "Direction[Right]",
}

var fromRuneMap = map[rune]Direction{
	'^': Up,
	'v': Down,
	'<': Left,
	'>': Right,
}

func (d Direction) String() string {

	ret, ok := stringMap[d]
	if !ok {
		log.Fatalf("invalid direction enum: %d", d)
	}

	return ret
}

func TryParseFromRune(r rune) (Direction, error) {

	ret, ok := fromRuneMap[r]
	if !ok {
		return Up, fmt.Errorf("invalid direction rune '%c'", r)
	}

	return ret, nil
}

func MustParseFromRune(r rune) Direction {

	ret, err := TryParseFromRune(r)
	if err != nil {
		log.Fatal(err)
	}

	return ret
}
