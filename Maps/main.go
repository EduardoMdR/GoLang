package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

var m1 = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	// Um map mapeia chaves para valores
	m = make(map[string]Vertex)
	// A função make retorna um map com um tipo determinado
	m["Bell Labs"] = Vertex{ // A chave "Bell labs" recebe os valores passados
		40.68433, -74.39967,
	}

	fmt.Println(m["Bell Labs"])

	fmt.Println(m1)
	fmt.Println(m1["Google"])

	// Mudando meu maps
	m2 := make(map[string]int)

	m2["Answer"] = 42 // Posso inserir elemento no meu map
	fmt.Println("The value:", m2["Answer"])

	m2["Answer"] = 48 // Posso alterar
	fmt.Println("The value:", m2["Answer"])

	delete(m2, "Answer") // Posso apagar
	fmt.Println("The value:", m2["Answer"])

	v, ok := m2["Answer"] // Testar se v está presente
	// v é um booleano
	fmt.Println("The value:", v, "Present?", ok)
}
