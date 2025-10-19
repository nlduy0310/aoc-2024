package direction

import (
	"slices"
)

type Direction int

const (
	North Direction = iota
	South
	West
	East
)

var clockwiseDirections = []Direction{
	North,
	East,
	South,
	West,
}

var namesMap = map[Direction]string{
	North: "North",
	South: "South",
	West:  "West",
	East:  "East",
}

var signsMap = map[Direction]string{
	North: "^",
	South: "v",
	West:  "<",
	East:  ">",
}

func (d Direction) TurnedLeft() Direction {
	clockwiseIdx := slices.Index(clockwiseDirections, d)
	if clockwiseIdx == -1 {
		invalidEnumPanic(d)
	}

	if clockwiseIdx == 0 {
		return clockwiseDirections[len(clockwiseDirections)-1]
	}

	return clockwiseDirections[clockwiseIdx-1]
}

func (d Direction) TurnedRight() Direction {
	clockwiseIdx := slices.Index(clockwiseDirections, d)
	if clockwiseIdx == -1 {
		invalidEnumPanic(d)
	}

	if clockwiseIdx == len(clockwiseDirections)-1 {
		return clockwiseDirections[0]
	}

	return clockwiseDirections[clockwiseIdx+1]
}

func (d Direction) Name() string {
	name, ok := namesMap[d]
	if !ok {
		invalidEnumPanic(d)
	}

	return name
}

func (d Direction) Sign() string {
	sign, ok := signsMap[d]
	if !ok {
		invalidEnumPanic(d)
	}

	return sign
}

func Directions() []Direction {
	return []Direction{North, South, West, East}
}
