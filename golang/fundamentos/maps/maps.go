package main

import "fmt"

func main() {

	pessoa := map[string]string{ // map[tipo_da_chave]valores_do_map
		"nome":  "Rodrigo",
		"idade": "40", // nao pode ser int pq o map foi criado para aceitar apenas string
	}

	fmt.Println(pessoa)
	fmt.Println("Para acessar um valor do map:", pessoa["nome"]) // nao permite pessoa.nome igual no struct

	// map aninhados

	pessoa2 := map[string]map[string]string{
		"nome": {
			"primeiro":  "Rodrigo",
			"sobrenome": "Martins",
		},
	}

	fmt.Println("Map com outro map nos valores", pessoa2)
	fmt.Println("para acessar os valores", pessoa2["nome"]["primeiro"])
	delete(pessoa2["nome"], "sobrenome")
	fmt.Println("removeu a chave sobrenome:", pessoa2)

	pessoa2["endereco"] = map[string]string{
		"logradouro": "Rua da Direita",
		"numero":     "22",
	}

	fmt.Println("adicionando chave endereco:", pessoa2)

	pessoa2["nome"]["sobrenome"] = "Martins"
	fmt.Println("adicionando novamente chave sobrenome ao map nome:", pessoa2)

}
