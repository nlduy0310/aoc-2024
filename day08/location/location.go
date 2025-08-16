package location

import "fmt"

type Location struct {
	Row int
	Col int
}

func NewLocation(row, col int) Location {

	// if row < 0 {
	// 	panic(fmt.Sprintf("a location's row index must not be negative, received %d", row))
	// }

	// if col < 0 {
	// 	panic(fmt.Sprintf("a location's column index must not be negative, received %d", col))
	// }

	return Location{
		Row: row,
		Col: col,
	}
}

func (l *Location) String() string {

	return fmt.Sprintf(
		"Location[row=%d, col=%d]",
		l.Row, l.Col,
	)
}
