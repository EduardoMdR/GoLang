package main

import (
	"fmt" // Biblioteca padrão
	"math"
	"time"
)

// var declara uma lista de variaveis
// var c, python, java
var i, j int = 1, 2

// Constantes são decladas com const
const nome = "Edu"

func main() {
	fmt.Println("Hello World")
	fmt.Println("Agora são", time.Now())

	// Letra maiuscula significa que é exportado de algum lugar -> Sqrt é exportado de math
	fmt.Printf("A raiz de 10 %g \n", math.Sqrt(10))
	// fmt.Println(math.Pi)

	// variavel sem tipo inferido é definida pelo valor recebido
	// Dentro de uma função eu posso utilizar := e var
	// Fora da função eu preciso de uma palavra chave (var, func)
	var c, python, java = true, false, "nop!"
	fmt.Println(i, j, c, python, java)

	// Variaveis declaradas sem um valor inicial receberão 'zero'
	var inteiro int
	var booleano bool
	var str string
	fmt.Printf("%v %v %q\n", inteiro, booleano, str)

	fmt.Println(nome)

	// A expressão do tipo T(v) converge o valor v para o tipo T

	// Ponteiros
	i, j := 42, 2701
	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)
	p = &j
	*p /= 37
	fmt.Println(j)
}
