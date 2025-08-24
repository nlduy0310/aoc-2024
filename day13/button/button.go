package button

import "fmt"

type Button struct {
	Name   string
	XShift int
	YShift int
	Cost   int
}

func NewButton(name string, xShift, yShift int, cost int) Button {

	return Button{
		Name:   name,
		XShift: xShift,
		YShift: yShift,
		Cost:   cost,
	}
}

func (b Button) String() string {

	return fmt.Sprintf("Button[name=%s, xShift=%d, yShift=%d, cost=%d]", b.Name, b.XShift, b.YShift, b.Cost)
}
