package main

import (
	"fmt"
	"reflect"
)

func main() {
	nome := "Rodrigo" // Tipo da variavel inferido // O var pode ser substituido por :=
	versao := 1.1     // Tipo da variavel inferido, o go cria como float64, // O var pode ser substituido por :=
	fmt.Println("Este programa está na versão", versao)
	fmt.Println("O tipo da variavel nome é ", reflect.TypeOf(nome))
	fmt.Println("O tipo da variavel versao é ", reflect.TypeOf(versao))
}
