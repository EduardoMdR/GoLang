package main

import (
	"fmt"
	"time"
)

// Programas Go expressam estados de erro com valores error.
type MeuErro struct {
	Quando   time.Time
	Mensagem string
}

// Semelhante ao Stringer, ou seja, quando retorna um erro
// sempre vai retornar com essa mensagem padrão escrita aqui
func (erro *MeuErro) Error() string {
	return fmt.Sprintf("at %v, %s",
		erro.Quando, erro.Mensagem)
}

// O método run() retorna um erro no formato:
// MeuErro{time.Now, "Não funcionou"}
func run() error {
	return &MeuErro{
		time.Now(),
		"Não funcionou",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

// type error interface {
// 	Error() string
// }
