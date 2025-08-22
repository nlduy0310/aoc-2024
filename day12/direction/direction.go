package direction

import "fmt"

type Direction int

const (
	North Direction = iota
	South
	West
	East
)

func (d Direction) String() string {

	switch d {
	case North:
		return "Direction[North]"
	case South:
		return "Direction[South]"
	case West:
		return "Direction[West]"
	case East:
		return "Direction[East]"
	default:
		panic(fmt.Sprintf("invalid direction enum %d", d))
	}
}
