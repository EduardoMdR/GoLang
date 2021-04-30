package main

import (
	"fmt"
)

type Interface interface {
	M()
}

type Tipo struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t Tipo) M() {
	fmt.Println(t.S)
}

func main() {
	var i Interface = Tipo{"hello"}
	i.M()

	// Por dentro, os valores de interface podem ser pensados
	// como uma tupla de um valor de um tipo concreto:
	// (valor, tipo)
	describe(i)

	i = &Tipo{"Hello"}
	describe(i)

	// Uma interface vazia pode conter valores de qualquer tipo.
	var i2 interface{}
	describe(i2)

	i2 = 42
	describe(i2)

	i2 = "hello"
	describe(i2)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
