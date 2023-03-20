package main

import (
	"fmt"
	"pacote/interno"

	"github.com/badoux/checkmail" // para instalar execute no terminal: go get github.com/badoux/checkmail
)

// exemplificar as regras de pacote do golang
// go run main.go
// go mod tidy -> limpar os imports nao usados e declarados no go.mod
func main() {
	fmt.Println("Escrevendo do main")
	escrever()
	interno.Escrever()

	fmt.Println(checkmail.ValidateFormat("123456"))
	fmt.Println(checkmail.ValidateFormat("rdgmart@email.com"))
}

// se funcao iniciar com letra minuscula
// será visível apenas dentro do mesmo pacote onde está os demais (similar ao protected/private de linguagens OO)
func escrever() {
	fmt.Println("escrevendo metodo privado do pacote main")
}
