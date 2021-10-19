package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main() {
	// Exemplo de matriz
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1], a)

	// Exemplo de Slice (s)
	primes := [6]int{2, 3, 5, 7, 11, 13}
	var primos []int = primes[1:4] // intervalo aberto (de 1 á 3)
	fmt.Println(primos)            // imprime 3, 5, 7

	// Slice são como referências para matrizes
	names := [4]string{ // poderia utilizar apenas []string
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	slice1 := names[0:2]
	slice2 := names[1:3]
	fmt.Println(slice1, slice2)

	slice2[0] = "XXX"
	fmt.Println(slice1, slice2)
	fmt.Println(names)

	// Toda Slice tem uma capacidade e um tamanho
	// capacidade: nº de elementos na matriz subjacente a partir do primeiro elemento na slice
	// tamanho: nº de elementos que ela contém

	s1 := []int{2, 3, 5, 7, 11, 13}
	printSlice(s1)

	// Slice the slice to give it zero length.
	s1 = s1[:0]
	printSlice(s1)

	// Extend its length.
	s1 = s1[:5] // s1 recebe s1 de 0 até 4
	printSlice(s1)

	// Drop its first two values.
	s1 = s1[2:] // s1 recebe s1 de 2 até acabar
	printSlice(s1)

	// o valor 0 de uma slice é nil
	var s []int
	fmt.Println(s, len(s), cap(s))
}
