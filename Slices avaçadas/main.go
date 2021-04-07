package main

import (
	"fmt"
	"strings"
)

func main() {
	// A função make aloca uma matriz zerada e retorna uma slice que se refere a essa matriz:
	a := make([]int, 5) // len(a)=5
	printSlice("a", a)

	b := make([]int, 0, 5) // cap(b) = 0, len(a)=5
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)

	// As slices podem conter qualquer tipo, incluindo outras slices.
	board := [][]string{
		{"1", "2", "3"}, // == []string{"1", "2", "3"}
		{"4", "5", "6"},
		{"7", "8", "9"},
	}
	board[1][1] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	// Append é uma função padrão de Go para adicionar elementos em uma slice
	var s2 []int
	fmt.Printf("len=%d cap=%d %v\n", len(s2), cap(s2), s2)

	// append works on nil slices.
	s2 = append(s2, 0)
	// o primeiro elemento é a slice e so segundo (ou mais) é o elemento que vou adicionar
	// Se a matriz anterior de s for muito pequena para caber todos os valores uma matriz gigante será alocada.
	// A slice retornada apontará para a nova matriz alocada.
	fmt.Printf("len=%d cap=%d %v\n", len(s2), cap(s2), s2)

	// The slice grows as needed.
	s2 = append(s2, 1)
	fmt.Printf("len=%d cap=%d %v\n", len(s2), cap(s2), s2)

	// We can add more than one element at a time.
	s2 = append(s2, 2, 3, 4)
	fmt.Printf("len=%d cap=%d %v\n", len(s2), cap(s2), s2)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
