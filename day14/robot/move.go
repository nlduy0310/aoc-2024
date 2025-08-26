package robot

type MoveOption = func(*Robot)

func ClampXAfterMove(xMin, xMax int) MoveOption {

	return func(r *Robot) {
		r.CurrentPosition.ClampX(xMin, xMax)
	}
}

func ClampYAfterMove(yMin, yMax int) MoveOption {

	return func(r *Robot) {
		r.CurrentPosition.ClampY(yMin, yMax)
	}
}

func (r *Robot) Move(options ...MoveOption) {

	r.CurrentPosition.Move(r.velocityX, r.velocityY)

	for _, opt := range options {
		opt(r)
	}
}
