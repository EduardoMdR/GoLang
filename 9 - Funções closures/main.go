package main

import "fmt"

// Um closure é uma função valor que referencia variáveis de fora de seu corpo

func contador() func(int) int {
	soma := 0
	return func(x int) int {
		soma += x
		return soma
	}
}

func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	positivo, negativo := contador(), contador()
	for i := 0; i < 10; i++ {
		fmt.Println(positivo(i), negativo(-2*i))
	}

	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}

	somador := 0
	// Um exemplo um pouco mais básico sobre closure
	counter := func() int {
		somador += 1
		return somador
	}

	fmt.Println(counter())
	fmt.Println(counter())
}
