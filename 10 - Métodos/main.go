package main

import (
	"fmt"
	"math"
)

type Dados struct {
	X, Y float64
}

// Go não tem classes. No entanto, você pode definir métodos em tipos.
// O método é uma função com um argumento receptor especial.
func (v Dados) Metodo1() float64 {
	// o método Abs1 tem um receptor do tipo Dados chamado v.
	return v.X + v.Y
} // Posso chamar como valor.Metodo1()

// Método é apenas uma função com um argumento receptor.
func Metodo2(v Dados) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
} // Posso chamar como Metodo2(valor)

type Float float64

// É possível declarar um método em tipos que não seja struct
func (f Float) Metodo3() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	valor1 := Dados{1, 2}
	fmt.Println(valor1.Metodo1())

	valor2 := Dados{3, 4}
	fmt.Println(Metodo2(valor2))

	valor3 := Float(-math.Sqrt2)
	fmt.Println(valor3.Metodo3())
}
