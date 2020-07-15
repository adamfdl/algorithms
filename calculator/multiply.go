package calculator

type multiplyCommand struct {
	A, B int
}

func (m multiplyCommand) calculate() int {
	return m.A * m.B
}
