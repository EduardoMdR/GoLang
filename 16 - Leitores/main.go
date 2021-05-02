package main

import (
	"fmt"
	"io"
	"strings"
)

// A interface io.Reader tem um método Read:
// func (T) Read(b []byte) (n int, err error)

func main() {
	texto := strings.NewReader("Hello, Reader!")
	fmt.Println(texto)

	b := make([]byte, 8)
	fmt.Println(b)

	// Read popula uma slice de bytes passados com dados e retorna o número
	// de bytes populados e um valor de erro. Este retorna um erro io.EOF
	// quando o fluxo termina.
	for {
		n, err := texto.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
