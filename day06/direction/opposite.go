package direction

var oppositeDirectionsMap map[Direction]Direction = map[Direction]Direction{
	Up:    Down,
	Down:  Up,
	Left:  Right,
	Right: Left,
}

func (d Direction) Opposite() Direction {

	ensureValid(d)

	return oppositeDirectionsMap[d]
}
