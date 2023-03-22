package interno

import "fmt"

// Escrever registra uma mensagem na tela
func Escrever() { // ao inicar metodo com letra maiuscula fica publica para outros pacotes invocar
	fmt.Println("Escrevendo do metodo interno")
	escreverInterno2() // consegue invocar pois interno2 foi declaro com o mesmo package -> interno

}
