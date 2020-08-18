package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

//Esta funcao eh a primeira funcao a ser executada pelo programa escrito em Go.
func main() {

	//Exibindo o menu de introducao
	exibirIntroducao()

	//Criando a lista de urls
	var listaDeUrlsDeSites []string

	//Executando o loop enquanto o comando for diferente de 0
	for {

		//Obtendo o comando do usuario
		exibirMenu()
		comando := lerComando()

		//Executando o comando do usuario
		switch comando {

		//Sair do programa
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)

		//Adicionando sites a lista
		case 1:
			adicionarSites(&listaDeUrlsDeSites)

		//Mostrando a lista de sites
		case 2:
			mostrarSites(&listaDeUrlsDeSites)

		//Removendo sites da lista
		case 3:
			removerSites(&listaDeUrlsDeSites)

		//Executando o monitoramento
		case 4:
			iniciarMonitoramento(&listaDeUrlsDeSites)

		//Mostrar os logs
		case 5:

		//Carregar lista a partir de um arquivo
		case 6:

		default:
			fmt.Println("Comando Inválido!")
			continue
		}

	}

}

//Esta funcao serve para exibir a mensagem de bem vindo ao programa e explica a utilidade do programa.
func exibirIntroducao() {

	fmt.Println("Bem vindo ao monitor de websites.")
	fmt.Println("Este programa serve para monitorar o status de determinados websites.")
}

//Esta funcao serve para obter o comando digitado pelo usuario apos exibir o menu do programa.
//Retorno: retorna um inteiro de 8 bits indicando qual comando foi solicitado pelo usuario.
func lerComando() int8 {
	var comando int8
	fmt.Scan(&comando)
	return comando
}

//Esta funcao serve para exibir os comandos que o usuario pode executar.
func exibirMenu() {
	fmt.Println("O que deseja fazer?")
	fmt.Println("0 - Sair do programa.")
	fmt.Println("1 - Adicionar sites a lista de sites monitorados.")
	fmt.Println("2 - Mostrar sites da lista de sites monitorados.")
	fmt.Println("3 - Remover sites da lista de sites monitorados.")
	fmt.Println("4 - Iniciar o monitoramento.")
}

//Esta funcao serve para realizar o monitoramento do(s) site(s).
//Parametro: lista: Ponteiro para a lista de url de sites que serao monitorados
func iniciarMonitoramento(lista *[]string) {
	fmt.Println("Monitorando...")

	for _, url := range *lista {
		monitorarSite(url)
	}
}

//Esta funcao serve para realizar o monitoramento de um site.
//Parametro: urlSite: Url do site a ser monitorado.
func monitorarSite(urlSite string) {
	resposta, erro := http.Get(urlSite)

	fmt.Println("URL:", urlSite)
	fmt.Println("Resposta:", resposta)
	fmt.Println("Erro", erro)
}

//Esta funcao serve para adicionar sites a lista de sites a serem monitorados
//Parametro: lista: Ponteiro para a lista de url de sites que serao monitorados
func adicionarSites(lista *[]string) {
	fmt.Println("Adicionando sites a lista...")
	adicionarSites := true
	for adicionarSites {

		//Lendo o endereco do site
		fmt.Print("URL do site: ")
		var url string
		fmt.Scan(&url)

		//Verificando se esta correto
		fmt.Print("O endereco do site: ", url, " está correto? ")
		var resposta string
		fmt.Scan(&resposta)
		resposta = strings.ToLower(resposta)
		resposta = string(resposta[0])
		if resposta != string('s') {
			continue
		} else {

			//Adicionando o site a lista
			*lista = append(*lista, url)

			//Verificando se o usuario deseja adicionar mais sites
			fmt.Print("Deseja adicionar mais sites a lista? ")
			var resposta string
			fmt.Scan(&resposta)
			resposta = strings.ToLower(resposta)
			resposta = string(resposta[0])
			if resposta != string('s') {
				adicionarSites = false
			}
		}
	}
}

//Esta funcao serve para mostrar todos os urls dos sites da lista
//Parametro: lista: Ponteiro para a lista de url de sites que serao monitorados
func mostrarSites(lista *[]string) {
	fmt.Println("Mostrando lista de sites...")
	fmt.Println(lista)
}

//Esta funcao serve para remover uma ou mais urls de sites da lista
//Prametro: lista: Ponteiro para a lista de url de sites que serao monitorados
func removerSites(lista *[]string) {
	fmt.Println("Removendo sites da lista...")
	if len(*lista) > 0 {

		//Mostrando a lista de sites
		for indice, url := range *lista {
			fmt.Println("[", indice+1, "]", url)
		}

		//Perguntando qual deles o usuario
		fmt.Print("Qual o numero da url deseja remover? ")
		var indice int
		fmt.Scan(&indice)
		indice--

		//Removendo a url da lista
		*lista = append((*lista)[:indice], (*lista)[(indice+1):]...)
	}
}
