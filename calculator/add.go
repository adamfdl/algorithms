package calculator

type addCommand struct {
	A, B int
}

func (a addCommand) calculate() int {
	return a.A + a.B
}
