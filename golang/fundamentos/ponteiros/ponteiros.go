package main

import "fmt"

func main() {
	var variavel1 int = 10
	var variavel2RecebendoPorCopia = variavel1

	fmt.Println("var1:", variavel1, "var2:", variavel2RecebendoPorCopia)

	variavel1++

	fmt.Println("var1:", variavel1, "var2:", variavel2RecebendoPorCopia) // nao altera o valor da variavel2RecebendoPorCopia

	// Por referência

	var variavel3RecebendoPorReferencia *int // o valor inicial de um ponteiro é <nil>
	variavel3RecebendoPorReferencia = &variavel1

	fmt.Println("var1:", variavel1, "var2:", variavel2RecebendoPorCopia,
		"var3 ponteiro:", variavel3RecebendoPorReferencia, "var3 valor:", *variavel3RecebendoPorReferencia) // * antes do nome da variavel é para acessar o valor do ponteiro

	variavel1++

	fmt.Println("var1:", variavel1, "var2:", variavel2RecebendoPorCopia, "var3:", *variavel3RecebendoPorReferencia)

}
