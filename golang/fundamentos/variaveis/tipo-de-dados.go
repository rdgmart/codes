package main

import (
	"errors"
	"fmt"
)

// Tipos basicos existente no GO
func main() {

	// temos 4 tipos de numero inteiros: int, int8, int16, int32, int64  (esses aceitam negativo)
	// o tipo int usa a arquitetura do computado para inferir
	// o tipo uint não aceita negativos e por isso dobra o maximo permitido

	var maxInt8 int8 = 127 // -128 to 127
	fmt.Println("Valor maximo para int8: ", maxInt8)

	var maxInt16 int16 = 32767 // -32768 to 32767
	fmt.Println("Valor maximo para int16: ", maxInt16)

	var maxInt32 int32 = 2147483647 // -2147483648 to 2147483647
	fmt.Println("Valor maximo para int32: ", maxInt32)

	var maxInt64 int64 = 9223372036854775807 // -9223372036854775808 to 9223372036854775807
	fmt.Println("Valor maximo para int64: ", maxInt64)

	var numeroComRune rune = 123456 // rune alias que representa int32
	fmt.Println("Valor declarado com rune ", numeroComRune)

	var numeroComByte byte = 127 // byte alias equivale a uint8
	fmt.Println("Valor declarado com rune ", numeroComByte)

	// float segue a mesma lógica do int mas temos apenas o float32, float64
	var numeroReal float32 = 12345.99
	fmt.Println("Valor declarado com float32 ", numeroReal)

	var numeroReal2 float64 = 1234599999.99
	fmt.Println("Valor declarado com float64 ", numeroReal2)

	// string ver o declaracao.go

	// char é um int e a curiosidade é que retorna o numero da representacao da tabela ASC
	char := 'A' // nao aceita mais de um caracter
	fmt.Println("Na tabela ASC a letra A é: ", char)

	// booleano return true ou false
	var verdadeiro bool = true
	fmt.Println("Valor declarado com bool ", verdadeiro)

	var naoIniciado bool // de nao for declarado por padrao é false
	fmt.Println("Valor declarado com bool e nao iniciado", naoIniciado)

	// error também é um tipo de dados em GO
	var erro error // qdo não iniciado retorno <nil>
	fmt.Println("Valor declarado com error ", erro)
	fmt.Println("Valor foi iniciado ", erro != nil)
	// para iniciar
	var erroIniciado error = errors.New("Erro declarado") // qdo não iniciado retorno <nil>
	fmt.Println("Valor declarado com error e iniciado: ", erroIniciado)
}
