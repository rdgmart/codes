package main

import "fmt"

func main() {

	var calculaSoma = func(n1, n2 int) int {
		return n1 + n2
	}

	fmt.Println("Resultado da funcao calcula soma 5+15=", calculaSoma(5, 15))

	resultSoma, resultSubtr := calculaSomaSubtracao(10, 9) // para omitir um dos retorno utilize: _, resultSubtr := calculaSomaSubtracao(10, 9)
	fmt.Println("Resultado da funcao calculaSomaSubtracao soma 10+9=", resultSoma)
	fmt.Println("Resultado da funcao calculaSomaSubtracao subtraca0 10-9=", resultSubtr)

}

func calculaSomaSubtracao(n1, n2 int) (int, int) {
	return (n1 + n2), (n1 - n2)
}
