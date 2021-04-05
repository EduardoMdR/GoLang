package main

import (
	"fmt"
)

func add(x int, y int) int {
	// Posso omitir valor de parametros consecutivos (x, y int)
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

func main() {
	fmt.Println(add(10, 20))

	a, b := swap("Mundo", "Olá")
	fmt.Println(a, b)

	fmt.Println(split(17))
}
