package main

import "fmt"

type Dados struct {
	X, Y int
}

var (
	v1       = Dados{1, 2}  // has type Vertex
	v2       = Dados{X: 1}  // Y:0 is implicit
	v3       = Dados{}      // X:0 and Y:0
	ponteiro = &Dados{1, 2} // has type *Vertex
)

func main() {
	fmt.Println(Dados{1, 2})
	v := Dados{1, 2}
	v.X = 0
	fmt.Println(v.X)

	p := &v
	p.X = 1e9 // isso é igula a (*p).X
	fmt.Println(v)

	fmt.Println(v1, ponteiro, v2, v3)

	// Quase uma Struct utilizando Slice
	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}
