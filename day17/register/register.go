package register

type Register struct {
	Name         string
	LiteralValue int
}

func NewEmptyRegister(name string) *Register {
	return &Register{
		Name:         name,
		LiteralValue: 0,
	}
}

func NewRegister(name string, initialValue int) *Register {
	return &Register{
		Name:         name,
		LiteralValue: initialValue,
	}
}
