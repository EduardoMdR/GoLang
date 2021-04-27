package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func elevado(x, y, lim float64) float64 {
	// posso começar o if com uma declaração
	if valor := math.Pow(x, y); valor < lim {
		return valor
	}
	return lim
}

func main() {
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

	// Defer adia a execusão da ação até o final da função
	defer fmt.Println("world")
	fmt.Println("Hello")

	// funciona como uma pilha, o ultimo a entrar é o primeiro a sair
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
}
