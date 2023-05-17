package main

import (
	"fmt"
	"strconv"
)

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
	fmt.Println("Estrutura de controle: Switch")
	fmt.Println(diaDaSemana(1))
	fmt.Println(diaDaSemana2(3))
	fmt.Println(exemploComFallThrough(13))

}

func diaDaSemana(numero int) string {

	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda"
	case 3:
		return "Terça-Feira"
	case 4:
		return "Quarta-Feira"
	case 5:
		return "Quinta-Feira"
	case 6:
		return "Sexta-Feira"
	case 7:
		return "Sábado"
	default:
		return "Esse dia da semana não existe"
	}
}

func diaDaSemana2(numero int) string { // nesse caso a condição fica no case
	var diaDaSemana string
	switch {
	case numero == 1:
		diaDaSemana = "Domingo"
	case numero == 2:
		diaDaSemana = "Segunda"
	case numero == 3:
		diaDaSemana = "Terça-Feira"
	case numero == 4:
		diaDaSemana = "Quarta-Feira"
	case numero == 5:
		diaDaSemana = "Quinta-Feira"
	case numero == 6:
		diaDaSemana = "Sexta-Feira"
	case numero == 7:
		diaDaSemana = "Sábado"
	default:
		diaDaSemana = "Esse dia da semana não existe"
	}

	return diaDaSemana
}

func exemploComFallThrough(numero int) string {
	var resposta string
	switch numero {
	case 13:
		resposta = "Esse não é o numero correto: " + strconv.Itoa(numero)
		fallthrough
	case 22:
		resposta = "O numero correto é: 22 ao inves desse: " + strconv.Itoa(numero)
	default:
		resposta = "O numero escolhido foi: " + strconv.Itoa(numero)
	}

	return resposta
}
