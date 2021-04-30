package main

import (
	"fmt"
)

type Dados struct {
	X, Y int
}

func (v Dados) Metodo() int {
	return v.X + v.Y
}

// É possivel declarar um método com ponteiro receptor
func (v *Dados) Escalar(f int) {
	v.X = v.X * f
	v.Y = v.Y * f
	// fmt.Println(v)
}

// Utilizando ponteiros eu altero o valor de 'Dados' na função main
// Utilizando valor eu altero sobre uma copia de 'Dados'

func main() {
	v := Dados{1, 2}
	v.Escalar(10) // É igual a (&v).Escalar(5)

	// fmt.Println(v)
	fmt.Println(v.Metodo())
}

// Da mesma maineira que se aplica ponteiros em métodos,
// são aplicados ponteiros nas funções

// func Escalar(v *Dados, f int) {
// 	v.X = v.X * f
// 	v.Y = v.Y * f
// }

// Escalar(&v, 10)
