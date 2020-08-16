package main

import (
	"fmt"
	"reflect"
)

func main() {

	//Atribuindo valores as variaveis
	nome := "Henrique"
	idade := 20
	versao := 1.1

	//Imprimindo os valores das variaveis
	fmt.Println("Olá", nome)
	fmt.Println("sua idade é", idade)
	fmt.Println("Este programa está na versão", versao)

	//Imprimindo os tipos das variaveis
	fmt.Println("O tipo da variável nome é:", reflect.TypeOf(nome))
	fmt.Println("O tipo da variável idade é:", reflect.TypeOf(idade))
	fmt.Println("O tipo da variável versao é:", reflect.TypeOf(versao))
}
