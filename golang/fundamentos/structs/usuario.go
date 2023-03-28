package main

import "fmt"

type pessoa struct {
	nome  string
	idade uint8 // somente positivos
}

type usuario struct {
	id       int
	pessoa   // isso seria o que representa 'herança' em GO
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {

	var usuario1 usuario
	usuario1.id = 123
	usuario1.nome = "Rodrigo Martins"
	usuario1.idade = 40

	end := endereco{"Rua da direita", 22}
	usuario1.endereco = end

	fmt.Println(usuario1.nome, usuario1)

	usuario2 := usuario{456, pessoa{"Antonio Toledo", 29}, endereco{"Rua que virou avenida", 13}} //  declarando de forma inline
	fmt.Println(
		usuario2.nome,         // pode invocar o atributo direto o atributo nome de pessoa
		usuario2.pessoa.idade, // ou atraves do type pessoa declarado na struct usuario
		usuario2)

	usuario3 := usuario{id: 789, pessoa: pessoa{nome: "Manoel Dantesco"}} // declarando apenas alguns atributos
	fmt.Println(usuario3.nome, usuario3)
}
