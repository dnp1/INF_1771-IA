package main

// SquareQueue ..
type SquareQueue []*Square

func (Q *SquareQueue) add(p *Square) {
	*Q = append(*Q, p)
}

func (Q *SquareQueue) get() *Square {
	var val = (*Q)[0]
	*Q = (*Q)[1:len(*Q)]
	return val
}
