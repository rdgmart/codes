package main

import "fmt"

func main() {
	fmt.Println("Estrutura de controle: if, else...")

	numero := 22

	fmt.Println("O numero é:", numero)

	// nao usamos parenteses a menos que tenha mais de uma condição para ser avaliada: cond1 && (cond2 || cond3)
	if numero > 13 {
		fmt.Println("Ele é maior que o 13")
	} else if numero == 13 {
		fmt.Println("Ele não é igual ao 13")
	} else {
		fmt.Println("Ele é menor que o 13")
	}

	// o if com init, a variavel só fica acessível no escopo
	if outroNumero := numero + 5; outroNumero > 0 {
		fmt.Println("outro numero foi iniciado com: ", outroNumero, " e é maior que zero")
	} else {
		fmt.Println("outro numero foi iniciado com: ", outroNumero, " e é menor ou igual à zero")
	}

}
