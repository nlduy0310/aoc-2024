package position

import "fmt"

type Position struct {
	Row, Col int
}

func (p Position) String() string {

	return fmt.Sprintf("Position[row=%d , col=%d]", p.Row, p.Col)
}

func (p Position) Up() Position {

	return Position{Row: p.Row - 1, Col: p.Col}
}

func (p Position) Down() Position {

	return Position{Row: p.Row + 1, Col: p.Col}
}

func (p Position) Left() Position {

	return Position{Row: p.Row, Col: p.Col - 1}
}

func (p Position) Right() Position {

	return Position{Row: p.Row, Col: p.Col + 1}
}
