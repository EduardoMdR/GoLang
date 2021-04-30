package main

import "fmt"

type Pessoa struct {
	Nome  string
	Idade int
}

// A Stringer é um tipo que pode descrever-se como uma string. O pacote fmt
// (e muitos outros) olham para essa interface para mostrar valores.
func (p Pessoa) String() string {
	return fmt.Sprintf("%v (%v anos)", p.Nome, p.Idade)
}

func main() {
	a := Pessoa{"Eduardo Marques", 20}
	b := Pessoa{"Seu Tião", 91}
	fmt.Println(a, b)
}

// type Stringer interface {
// 	String() string
// }
