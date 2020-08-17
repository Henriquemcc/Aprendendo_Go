package main

import (
	"fmt"
	"os"
)

//A funcao main eh a primeira funcao a ser executada pelo programa escrito em Go.
func main() {

	exibirIntroducao()
	exibirMenu()

	comando := lerComando()

	switch comando {
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo os logs...")
	case 0:
		fmt.Println("Saindo do programa...")
		os.Exit(0)
	default:
		fmt.Println("Comando Inválido!")
		os.Exit(-1)
	}

}

//A funcao exibirIntroducao exibe a mensagem de bem vindo ao programa e explica a utilidade do programa.
func exibirIntroducao() {

	fmt.Println("Bem vindo ao monitor de websites.")
	fmt.Println("Este programa serve para monitorar o status de determinados websites.")
}

//A funcao lerComando obtem o comando digitado pelo usuario apos exibir o menu do programa.
//Esta funcao retorna um inteiro de 8 bits indicando qual comando foi solicitado pelo usuario.
func lerComando() int8 {
	var comando int8
	fmt.Scan(&comando)
	return comando
}

//A funcao exibirMenu exibe os comandos que o usuario pode executar.
func exibirMenu() {
	fmt.Println("O que deseja fazer: ")
	fmt.Println("0 - Sair do programa")
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir os logs")
}
