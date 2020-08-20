package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

var monitoramentos = 10
var delay = 5 * time.Minute
var nomeArquivoUrls = "urlsMonitoradas.txt"
var nomeArquivoLogs = "logs.txt"

//Esta funcao eh a primeira funcao a ser executada pelo programa escrito em Go.
func main() {

	//Exibindo o menu de introducao
	exibirIntroducao()

	//Criando a lista de urls
	listaDeUrlsDeSites := lerUrlsDoArquivo()

	//Executando o loop enquanto o comando for diferente de 0
	continuarloop := true
	for continuarloop {

		//Obtendo o comando do usuario
		exibirMenu()
		comando := lerComando()
		fmt.Println("Comando a ser executado:", comando)

		//Executando o comando do usuario
		switch comando {

		//Sair do programa
		case 0:
			fmt.Println("Saindo do programa...")
			continuarloop = false

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

		//Alterar configuracoes de monitoramento
		case 5:
			alterarConfiguracoes()

		//Mostrar os logs
		case 6:

		//Carregar lista a partir de um arquivo
		case 7:

		default:
			fmt.Println("Comando Inválido!")
		}

		fmt.Println("")

	}

	//Salvando as urls no arquivo
	salvarUrlsNoArquivo(listaDeUrlsDeSites)

}

//Esta funcao serve para exibir a mensagem de bem vindo ao programa e explica a utilidade do programa.
func exibirIntroducao() {

	fmt.Println("Bem vindo ao monitor de websites.")
	fmt.Println("Este programa serve para monitorar o status de determinados websites.")
}

//Esta funcao serve para obter o comando digitado pelo usuario apos exibir o menu do programa.
//Retorno: retorna um inteiro de 8 bits indicando qual comando foi solicitado pelo usuario.
func lerComando() int8 {
	var comando int8 = -1
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
	fmt.Println("5 - Alterar configurações de monitoramento.")
}

//Esta funcao serve para realizar o monitoramento do(s) site(s).
//Parametro: lista: Ponteiro para a lista de url de sites que serao monitorados
func iniciarMonitoramento(lista *[]string) {
	fmt.Println("Monitorando...")

	//Rodando testando os sites de acordo com o numero de vezes que serao testados
	for indice := 0; indice < monitoramentos; indice++ {

		//Imprimindo mensagem informando quanto de quanto foi executado
		fmt.Println("Teste", indice+1, "de", monitoramentos)

		//Realizando teste para cada site da lista
		for _, url := range *lista {
			monitorarSite(url)
			fmt.Println("")
		}

		//Dando delay nos testes
		time.Sleep(delay)
		fmt.Println("")
	}
}

//Esta funcao serve para realizar o monitoramento de um site.
//Parametro: urlSite: Url do site a ser monitorado.
func monitorarSite(urlSite string) {

	//Realizando o get do site
	resposta, erro := http.Get(urlSite)

	//Verificando se nao ocorreu nenhum erro ao realizar o get
	if erro != nil {
		fmt.Println("Um erro ocorreu:", erro)
		return
	}

	//Imprimindo mensagem sobre qual site esta sendo testado
	fmt.Println("Testando site:", urlSite)

	//Verificando se o site esta online
	online := resposta.StatusCode == 200

	//Imprimindo mensagem de sucesso ou erro de acordo com o estado do site
	if online {
		fmt.Println("O site", urlSite, "testado com sucesso!")
	} else {
		fmt.Println("Não foi possivel acessar o site", urlSite)
	}

	//Imprimindo a resposta do site
	fmt.Println("Resposta do site:")
	fmt.Println("-------------------------")
	fmt.Println(resposta)
	fmt.Println("-------------------------")

	//Registrando o status do site no arquivo log
	registrarLog(urlSite, online, time.Now())
}

//Esta funcao serve para adicionar sites a lista de sites a serem monitorados
//Parametro: lista: Ponteiro para a lista de url de sites que serao monitorados
func adicionarSites(lista *[]string) {

	//Exibindo mensagem de inicio do caso de uso
	fmt.Println("Adicionando sites a lista...")

	//Adicionando sites enquanto o usuario digitar sim
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

		//Permitindo ao usuario re-digitar o site caso ele esteja errado
		if resposta != string('s') {
			continue

			//Adicionando o site a lista caso ele esteja correto
		} else {

			//Adicionando o site a lista
			*lista = append(*lista, url)

			//Verificando se o usuario deseja adicionar mais sites
			fmt.Print("Deseja adicionar mais sites a lista? ")
			var resposta string
			fmt.Scan(&resposta)
			resposta = strings.ToLower(resposta)
			resposta = string(resposta[0])

			//Terminando o loop caso o usuario nao deseje mais adicionar mais sites
			if resposta != string('s') {
				adicionarSites = false
			}
		}
	}
}

//Esta funcao serve para mostrar todos os urls dos sites da lista
//Parametro: lista: Ponteiro para a lista de url de sites que serao monitorados
func mostrarSites(lista *[]string) {

	//Exibindo mensagem de inicio de caso de uso
	fmt.Println("Mostrando lista de sites...")

	//Imprimindo a lista
	fmt.Println(lista)
}

//Esta funcao serve para remover uma ou mais urls de sites da lista
//Prametro: lista: Ponteiro para a lista de url de sites que serao monitorados
func removerSites(lista *[]string) {

	//Exibindo mensagem de inicio de caso de uso
	fmt.Println("Removendo sites da lista...")

	//Verificando se eh possivel remover algum site
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

//Esta funcao serve para alterar as configuracoes de monitoramento
func alterarConfiguracoes() {

	//Exibindo mensagem de inicio de caso de uso
	fmt.Println("Alterando configurações de monitoramento...")

	//Obtendo o tempo de delay
	fmt.Println("Tempo de delay: ")

	//Valor do tempo de delay
	var valor int = -1
	for valor < 0 {

		//Obtendo o valor
		fmt.Print("Valor: ")
		fmt.Scan(&valor)

		//Imprimindo mensagem de valor invalido
		if valor < 0 {
			fmt.Println("Valor inválido!. Tente novamente.")
		}
	}

	//Unidade do tempo de delay
	delay = time.Duration(0)
	unidadeDeTempoDefinida := false
	for !unidadeDeTempoDefinida {
		var unidadeTempo string
		fmt.Print("Unidade de tempo (nanosegundos/microsegundos/milisegundos/segundos/minutos/horas) : ")
		fmt.Scan(&unidadeTempo)
		unidadeTempo = strings.ToLower(unidadeTempo)
		unidadeTempo = string(unidadeTempo[:3])

		//Calculando a unidade de tempo
		if unidadeTempo == "nan" {
			delay = time.Nanosecond * time.Duration(valor)
			unidadeDeTempoDefinida = true
		} else if unidadeTempo == "mic" {
			delay = time.Microsecond * time.Duration(valor)
			unidadeDeTempoDefinida = true
		} else if unidadeTempo == "mil" {
			delay = time.Millisecond * time.Duration(valor)
			unidadeDeTempoDefinida = true
		} else if unidadeTempo == "seg" {
			delay = time.Second * time.Duration(valor)
			unidadeDeTempoDefinida = true
		} else if unidadeTempo == "min" {
			delay = time.Minute * time.Duration(valor)
			unidadeDeTempoDefinida = true
		} else if unidadeTempo == "hor" {
			delay = time.Hour * time.Duration(valor)
			unidadeDeTempoDefinida = true
		} else {
			fmt.Println("Unidade de tempo invalida! Tente novamente.")
		}
	}

	//Obtendo a Quantidade de monitoramentos que deverao ser feitos
	monitoramentos = 0
	for monitoramentos <= 0 {
		fmt.Print("Quantidade de vezes que o(s) site(s) deve(m) ser monitorado(s): ")
		fmt.Scan(&monitoramentos)

		//Imprimindo mensagem de valor invalido
		if monitoramentos <= 0 {
			fmt.Println("Valor inválido! Tente novamente.")
		}
	}

}

//Esta funcao serve para ler as urls de um arquivo
//Retorno: Um ponteiro para uma slice de strings
func lerUrlsDoArquivo() []string {

	//Criando a variavel de retorno
	var urls []string

	//Criando o arquivo
	arquivo, erro := os.OpenFile(nomeArquivoUrls, os.O_RDONLY|os.O_CREATE, 0666)

	//Verificando se nenhum erro ocorreu ao abrir o arquivo
	if erro != nil {
		fmt.Println("Um erro ocorreu:", erro)
		return nil
	}

	//Criando um objeto para ler o arquivo
	leitor := bufio.NewReader(arquivo)

	//Lendo todo o conteudo do arquivo
	for erro != io.EOF {

		//Lendo uma linha do arquivo
		var url string
		url, erro = leitor.ReadString('\n')

		//Removendo os espacos da linha e adicionando na slice
		urls = append(urls, strings.TrimSpace(url))
	}

	//Fechando o arquivo
	arquivo.Close()

	//Retornando a lista das urls
	return urls
}

//Esta funcao serve para salvar as urls em um arquivo.
//Parametro: lista: Lista de urls que serao salvas no arquivo.
func salvarUrlsNoArquivo(lista []string) {

	//Abrindo o arquivo
	arquivo, erro := os.OpenFile(nomeArquivoUrls, os.O_WRONLY|os.O_CREATE, 0666)

	//Verificando se houve erro ao abrir o arquivo
	if erro != nil {
		fmt.Println("Um erro ocorreu:", erro)
		return
	}

	//Salvando as urls no arquivo
	for _, url := range lista {

		//Salvando a url no arquivo
		_, erro = arquivo.WriteString(url + "\n")

		//Verificando se algum erro ocorreu
		if erro != nil {
			fmt.Println("Um erro ocorreu:", erro)
		}
	}

	//Fechando o arquivo
	arquivo.Close()

}

//Esta funcao serve para salvar o log do status de um site no arquivo
//Parametro: site: Url do site que foi testado
//Parametro: status: Status do site: true = online, false = offline.
//Parametro: horario: Horario em que o site foi testado
func registrarLog(site string, status bool, horario time.Time) {

	//Abrindo o arquivo
	arquivo, erro := os.OpenFile(nomeArquivoLogs, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	//Verificando se houve erro ao abrir o arquivo
	if erro != nil {
		fmt.Println("Um erro ocorreu:", erro)
		return
	}

	//Escrevendo no arquivo de acordo com o status do site
	if status {
		_, erro = arquivo.WriteString(horario.String() + "\t-\t" + site + "\t\tonline\n")
	} else {
		_, erro = arquivo.WriteString(horario.String() + "\t-\t" + site + "\t\toffline\n")
	}

	//Verificando se houve erro ao escrever
	if erro != nil {
		fmt.Println("Um erro ocorreu:", erro)
	}

	//Fechando o arquivo
	arquivo.Close()
}
