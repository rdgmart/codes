package main

import "fmt"

func main() {

	fmt.Println("-------ATRIBUICAO-------")
	fmt.Println(" use = OU := ")            //
	var texto string = "exemplo atribuicao" // precisa declarar o tipo
	fmt.Println(texto, " com = ")           //

	fmt.Println("-------ARITIMETICOS-------") // devem ter o mesmo tipo. Por exemplo int8 + int32 nao funciona

	soma := 1 + 2 // nesse o tipo é inferido
	fmt.Println("Resultado soma (1+2): ", soma)

	sub := 50 - 9
	fmt.Println("Resultado subtracao (50-9): ", sub)

	multiplicacao := 10 * 5
	fmt.Println("Resultado multiplicaco (10*5): ", multiplicacao)

	divisao := 10 / 2
	fmt.Println("Resultado divisao (10/2): ", divisao)

	restoDivisao := 10 % 3
	fmt.Println("Resultado resto da divisao (10-3): ", restoDivisao)

	fmt.Println("-------RELACIONAIS-------")
	fmt.Println(" 1 == 2", 1 == 2)
	fmt.Println(" 1 != 2", 1 != 2)
	fmt.Println(" 1 > 2", 1 > 2)
	fmt.Println(" 1 >= 2", 1 >= 2)
	fmt.Println(" 1 < 2", 1 < 2)
	fmt.Println(" 1 <= 2", 1 > 2)

	fmt.Println("-------LOGICOS-------")
	verdadeiro, falso := true, false
	fmt.Println("verdadeiro && falso =", verdadeiro && falso)
	fmt.Println("verdadeiro || falso =", verdadeiro || falso)
	fmt.Println("!verdadeiro =", !verdadeiro)
	fmt.Println("!falso =", !falso)

	fmt.Println("-------UNARIOS-------")
	numero := 10
	numero++
	fmt.Println("numero 10 aplicado o numero++ fica com o valor ", numero)
	numero -= 5 // numero = numero - 5
	fmt.Println("resultado anterior aplicado numero -= 5 fica com o valor ", numero)
	numero *= 2 // numero = numero * 2
	fmt.Println("resultado anterior aplicado numero *= 2  fica com o valor ", numero)

	fmt.Println("-------TERNARIO NAO TEM-------")
	fmt.Println("nao consegue fazer algo: numero > 5 ? 'maior' : 'menor' ")
	fmt.Println("nesse caso usar if/else")

}
