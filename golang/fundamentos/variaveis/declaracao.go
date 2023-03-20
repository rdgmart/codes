package main

import "fmt"

// Golang é fortemente tipado
// Toda variavel declarada deve ser usada
func main() {

	var variavel string = "variavel 1"
	fmt.Println(variavel)

	variavel2 := "variavel 2" // por inferencia: nesse caso declara e já atribui e por isso identifica o tipo. Deve usar :=
	fmt.Println(variavel2)

	var (
		variavel3 string = "declarando variavel 3..."
		variavel4 string = "e variavel 4 juntas."
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "variavel 5 ...", "e 6 declaradas e iniciadas juntas por inferência."

	fmt.Println(variavel5, variavel6)
}
