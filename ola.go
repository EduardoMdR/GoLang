package main

import (
	"fmt" 	// Biblioteca padrão
	"time"
	"math"
)

func add(x int, y int) int{
	// P0osso omitir valor de parametros consecutivos (x, y int)
	return x + y
}

func swap(x, y string) (string, string){
	// Posso retornar quantos resultados eu quero
	return y, x
}

func split(sum int) (x, y int){
	// A declaração retunr sem argumentos retorna os valores atuais dos resultados
	x = sum * 4 / 9
	y = sum - x
	return
}

// var declara uma lista de variaveis
// var c, python, java
var i, j int = 1, 2

// Constantes são decladas com const
const nome = "Edu"

func main(){
	fmt.Println("Hello World")
	fmt.Println("Agora são", time.Now())

	// Letra maiuscula significa que é exportado de algum lugar -> Sqrt é exportado de math
	fmt.Printf("A raiz de 10 %g \n", math.Sqrt(10))
	// fmt.Println(math.Pi)

	fmt.Println(add(10,20))

	a, b := swap("Mundo", "Olá")
	// variavel sem tipo inferido é definida pelo valor recebido
	fmt.Println(a, b)

	fmt.Println(split(17))

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
}
// A expressão do tipo T(v) converge o valor v para o tipo T


// Lista Fibonacci
// func fib() func() int{
// 	a, b := 0,1
// 	return func() int{
// 		a, b = b, a + b
// 		return a
// 	}
// }

// func main(){
// 	f := fib()
// 	fmt.Println(f(), f(), f(), f(), f(), f())
// }