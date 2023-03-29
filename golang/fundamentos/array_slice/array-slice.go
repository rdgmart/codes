package main

import (
	"fmt"
	"reflect"
)

func main() {

	var array1 [5]string    // o array tem tamanho fixo
	array1[2] = "Posicao 3" // consigo atribuir valor em posicao especifica

	fmt.Println(array1)

	array2 := [3]string{"Posicao 1", "Posicao 2", "Posicao 3"} // outra forma de declarar inicializando
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4} // ... iniciando dinamicamente, mas o tamanho depois será fixo
	fmt.Println(array3)            //  no array nao é permitido adicionar valor depois de criado

	slice := []int{5, 6, 7, 8, 9} // slice cria um ponteiro e aponta para a posicao de memoria do array criado
	fmt.Println(slice)
	slice = append(slice, 10) // consigo adicionar valor... no caso cria outro array e aponta para ele.
	fmt.Println(slice)

	fmt.Println("isso é um array:", reflect.TypeOf(array3))
	fmt.Println("isso é um slice:", reflect.TypeOf(slice))

	slice2 := array3[2:4]
	fmt.Println("slice criado com um pedaco do array: ", slice2)

	array3[2] = 99 // o slice tb aponta para essa posição... será que vai alterar para ele tb????

	fmt.Println("como o slice foi criado com um pedaco do array... ele tb sofre alteracao: ", slice2) // lembre que slice é criado referenciando um array...como um ponteiro

	// array interno

	// make cria um array com o tamanho maximo, e retorna um pedaco (slice) de acordo com o tamanho inicial
	slice3 := make([]float32, 10, 12) // tipo, tamanho inicial, tamanho maximo (se nao for especificado fica igual ao tamanho inicial)
	fmt.Println("\nslice criado com make:", slice3)
	fmt.Println("tamanho: ", len(slice3))
	fmt.Println("capacidade: ", cap(slice3))

	// o GO qdo atinge a capacidade ele criar outro array com o dobro do tamanho
	slice3 = append(slice3, 11)
	slice3 = append(slice3, 12)
	fmt.Println("\nslice adicionado mais 2 valores:", slice3)
	fmt.Println("tamanho: ", len(slice3))
	fmt.Println("capacidade: ", cap(slice3), "(antes de estourar)")
	slice3 = append(slice3, 13)
	fmt.Println("\nslice ultrapassando a capacidade:", slice3)
	fmt.Println("tamanho: ", len(slice3))
	fmt.Println("capacidade: ", cap(slice3), "(depois de estourar... dobrou)")
	fmt.Println("por isso o slice não tem tamanho limite ")

}
