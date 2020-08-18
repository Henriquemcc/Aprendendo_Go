package main

import (
	"fmt"
	"net/http"
	"os"
)

//A funcao main eh a primeira funcao a ser executada pelo programa escrito em Go.
func main() {

	exibirIntroducao()

	//Executando o loop enquanto o comando for diferente de 0
	for {

		//Obtendo o comando do usuario
		exibirMenu()
		comando := lerComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo os logs...")
		case 0:
			fmt.Println("Saindo do programa...")
		default:
			fmt.Println("Comando Inválido!")
			os.Exit(-1)
		}

		//Finalizando o loop
		if comando == 0 {
			break
		}

	}

	//Finalizando o programa
	os.Exit(0)

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

//A funcao iniciarMonitoramento realiza o monitoramento do(s) site(s).
func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	site := "https://henriquemcc101508968.wordpress.com"
	resposta, erro := http.Get(site)

	fmt.Println("Resposta:", resposta)
	fmt.Println("Erro", erro)

	if resposta.StatusCode == 200 {
		fmt.Println("O site", site, "foi carregado com sucesso")
	} else {
		fmt.Println("O site", site, "está com problemas. Status code:", resposta.StatusCode)
	}

}
