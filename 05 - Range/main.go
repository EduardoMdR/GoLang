package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	// O range itera minha variável e retorna seu valor e seu indice
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	// Caso queira utilizar so os indices
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}

	// Caso queira utilizar só os valores
	for _, value := range pow {
		fmt.Printf("%d, ", value)
	}
}
