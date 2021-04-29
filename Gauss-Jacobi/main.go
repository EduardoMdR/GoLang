package main

type Jacobi struct {
	X1, X2, X3, Y int
}

var (
	e1 = Jacobi{-5, 2, 2, 2}
	e2 = Jacobi{-2, 4, 1, -1}
	e3 = Jacobi{1, -2, 6, 1}
)

func main() {
	x1 := 0.0
	for i := 0; i < 10; i++ {
		x1 = (1 / e1.X1 * (e1.Y - e1.X2 - e1.X3))
	}
}
