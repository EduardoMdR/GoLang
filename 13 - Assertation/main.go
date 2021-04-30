package main

import "fmt"

func main() {
	var i interface{} = "hello"

	// type assertation fornece acesso ao valor de interface.

	// Se i detém uma 'string', então 'value' será o
	// valor subjacente e ok será true.
	value1, ok := i.(string)
	fmt.Println(value1, ok)

	value2, ok := i.(float64)
	fmt.Println(value2, ok)

	do(21)
	do("hello")
	do(true)
}

// Um type switch é uma construção que permite diversas
// type assertations em série
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("O dobro de %v é %v\n", v, v*2)
	case string:
		fmt.Printf("%q tem o tamanho de %v bytes\n", v, len(v))
	default:
		fmt.Printf("Que tipo é esse? n sei! %T!\n", v)
	}
}
