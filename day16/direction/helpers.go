package direction

import "fmt"

func invalidEnumPanic(d Direction) {
	panic(fmt.Sprintf("invalid direction enum value: %d", d))
}
