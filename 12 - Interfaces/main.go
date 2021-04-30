package main

import (
	"fmt"
	"math"
)

// Um tipo de interface é definida por um conjunto de métodos.
// Um valor de tipo de interface pode conter qualquer valor que implementa esses métodos.
type Interface interface {
	Metodo() float64
}

func main() {
	var a Interface
	f := Float(-math.Sqrt2)
	v := Dados{3, 4}

	a = f // a Float implementa Interface
	// fmt.Println(a.Metodo())
	a = &v // a *Dados implementa Interface

	fmt.Println(a.Metodo())
}

type Float float64

func (f Float) Metodo() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Dados struct {
	X, Y float64
}

func (v *Dados) Metodo() float64 {
	return v.X + v.Y
}
