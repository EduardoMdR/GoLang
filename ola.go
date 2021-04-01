package main

import (
	"fmt" // Biblioteca padrão
	"math"
	"runtime"
	"time"
)

func add(x int, y int) int {
	// P0osso omitir valor de parametros consecutivos (x, y int)
	return x + y
}

func swap(x, y string) (string, string) {
	// Posso retornar quantos resultados eu quero
	return y, x
}

func split(sum int) (x, y int) {
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

func elevado(x, y, lim float64) float64 {
	// posso começar o if com uma declaração
	if valor := math.Pow(x, y); valor < lim {
		return valor
	}
	return lim
}

func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	fmt.Println("Hello World")
	fmt.Println("Agora são", time.Now())

	// Letra maiuscula significa que é exportado de algum lugar -> Sqrt é exportado de math
	fmt.Printf("A raiz de 10 %g \n", math.Sqrt(10))
	// fmt.Println(math.Pi)

	fmt.Println(add(10, 20))

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

	sum := 0
	// a instrução inicial e pós-instrução são opcionais for ;i<10; {} ou ainda for i<10 {} == while
	for i := 0; i < 10; i++ {
		// i só está disponivel no escopo da declaração for
		sum += 1
	}
	fmt.Println(sum)

	fmt.Println(elevado(2, 3, 10), elevado(3, 3, 20))

	// Switch funciona como nas outras linguagens, mas não precisa do break
	switch so := runtime.GOOS; so {
	case "darwin":
		fmt.Println("Bom dia")
	case "linux":
		fmt.Println("Tô no Linux")
	default:
		fmt.Printf("%s \n", so)
	}

	tempo := time.Now()
	// Switch sem condição é o mesmo que switch true, e para na primeira verdade, ou no default
	switch {
	case tempo.Hour() < 12:
		fmt.Println("Bom dia")
	case tempo.Hour() < 17:
		fmt.Println("Boa tarde")
	default:
		fmt.Println("Boa noite")
	}

	f := fib()
	for i := 0; i < 10; i++ {
		fmt.Print(f())
	}

	// Defer adia a execusão da ação até o final da função
	defer fmt.Println("world")
	fmt.Println("Hello")

	// funciona como uma pilha, o ultimo a entrar é o primeiro a sair
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
}

// A expressão do tipo T(v) converge o valor v para o tipo T
