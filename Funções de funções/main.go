package main

import (
	"fmt"
	"math"
)

// Funções são valores também. Elas podem ser passadas assim como outros valores.
// Funções valores podem ser usadas como argumentos de funções e retornar valores.

func dados(passaDados func(float64, float64) float64) float64 {
	return passaDados(3, 4)
	// A função dados recebe '(passaDados func(float64, float64) float64)' como valor de entrada e retorna 'float64'
	// A função passaDados recebe '(3, 4)' e retorna 'float64'
}

func main() {
	hipotenusa := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hipotenusa(3, 4))

	fmt.Println(dados(hipotenusa))
	fmt.Println(dados(math.Pow))
}

// func dados(float64, float64) (float64, float64) {
// 	return 3, 4
// 	// Não consigo utilizar dados(hipotenusa ou math.Pow) dessa maneira
// }
