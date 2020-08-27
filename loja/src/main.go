package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"loja/produtos"
	"net/http"
	"os"
	"strings"
	"text/template"

	_ "github.com/lib/pq"
)

var nomeArquivoCredencialBancoDeDados = "CredencialBancoDeDados.env"

//Criando o template da aplicacao web
var templateDaAplicacaoWeb = template.Must(template.ParseGlob("templates/*.html"))

//Esta funcao eh a funcao principal, onde o programa comeca a ser executado.
func main() {
	bd := ConectarComBancoDeDados(ObterCredenciaisDeAcessoAoBancoDeDados(nomeArquivoCredencialBancoDeDados))
	defer bd.Close()
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

//Esta funcao serve para atender a requisicao '/'
//Parametro: w: Instancia da interface usada pelo manipulador http para construir a resposta http.
//Parametro: r: Instancia da struct request, que representa uma requisicao recebida pelo servidor ou enviada pelo cliente.
func index(w http.ResponseWriter, r *http.Request) {

	//Criando uma lista de produtos
	var listaDeProdutos []produtos.Produto

	//Criando o primeiro produto e adicionando a lista
	camiseta := produtos.Produto{}
	camiseta.SetNome("Camiseta")
	camiseta.SetDescricao("Azul")
	camiseta.SetPreco(20.00)
	camiseta.SetQuantidade(5)
	listaDeProdutos = append(listaDeProdutos, camiseta)

	//Criando o segundo produto e adicionando a lista
	computadorGamer := produtos.Produto{}
	computadorGamer.SetNome("Computador Gamer")
	computadorGamer.SetDescricao("Computador Gamer Poderoso")
	computadorGamer.SetPreco(30000.00)
	computadorGamer.SetQuantidade(5)
	listaDeProdutos = append(listaDeProdutos, computadorGamer)

	//Criando o terceiro produto e adicionando a lista
	videoGame := produtos.Produto{}
	videoGame.SetNome("Video game")
	videoGame.SetDescricao("Console de video game com 2TB de SSD.")
	videoGame.SetPreco(8000.00)
	videoGame.SetQuantidade(8)
	listaDeProdutos = append(listaDeProdutos, videoGame)

	//Criando o quarto produto e adicionando a lista
	smartphone := produtos.Produto{}
	smartphone.SetNome("Smartphone")
	smartphone.SetDescricao("Smartphone com Android 11.")
	smartphone.SetPreco(5000.00)
	smartphone.SetQuantidade(12)
	listaDeProdutos = append(listaDeProdutos, smartphone)

	//Criando o quinto produto e adicionando a lista
	roteador := produtos.Produto{}
	roteador.SetNome("Roteador WiFi")
	roteador.SetDescricao("Roteador Wifi Mesh Dual Band 2,4Ghz e 5Ghz")
	roteador.SetPreco(2000.0)
	roteador.SetQuantidade(6)
	listaDeProdutos = append(listaDeProdutos, roteador)

	//Executando aplicacao web
	templateDaAplicacaoWeb.ExecuteTemplate(w, "Index", listaDeProdutos)
}

//ObterCredenciaisDeAcessoAoBancoDeDados serve para obter as credenciais de acesso ao banco de dados de um arquivo.
//Parametro: nomeArquivo: Nome do arquivo em que sera lido para obter as credenciais de acesso ao banco de dados.
//Retorno: string: Nome de usuario do banco de dados.
//Retorno: string: Nome do banco de dados.
//Retorno: string: Senha do banco de dados
//Retorno: string: Host do banco de dados.
//Retorno: string: Modo de configuracao do protocolo SSL do banco de dados.
func ObterCredenciaisDeAcessoAoBancoDeDados(nomeArquivo string) (string, string, string, string, string) {
	//Abrindo o arquivo que contem as credenciais de acesso ao banco de dados Postgre SQL
	arquivo, erro := os.Open(nomeArquivo)

	//Exibindo mensagem de erro e abortando programa caso nao seja possivel ler o arquivo de credenciais do banco de dados
	if erro != nil {
		fmt.Println("Um erro ocorreu:", erro)
		os.Exit(-1)
	}

	//Criando um novo objeto para ler o arquivo
	leitor := bufio.NewReader(arquivo)

	//Lendo o nome de usuario
	nomeUsuario, erro := leitor.ReadString('\n')
	nomeUsuario = strings.TrimSpace(nomeUsuario)

	//Lendo o nome do banco de dados
	nomeBancoDeDados, erro := leitor.ReadString('\n')
	nomeBancoDeDados = strings.TrimSpace(nomeBancoDeDados)

	//Lendo a senha do banco de dados
	senha, erro := leitor.ReadString('\n')
	senha = strings.TrimSpace(senha)

	//Lendo o host do banco de dados
	host, erro := leitor.ReadString('\n')
	host = strings.TrimSpace(host)

	//Lendo o modo ssl do banco de dados
	ssl, erro := leitor.ReadString('\n')
	ssl = strings.TrimSpace(ssl)

	//Fechando o arquivo
	arquivo.Close()

	return nomeUsuario, nomeBancoDeDados, senha, host, ssl
}

//ConectarComBancoDeDados serve para abrir uma conexao com o banco de dados.
//Parametro: nomeUsuario: Nome de usuario do banco de dados.
//Parametro: nomeBancoDeDados: Nome do banco de dados.
//Parametro: senha: Senha do banco de dados.
//Parametro: host: Host do banco de dados.
//Parametro: ssl: Modo de configuracao do protocolo SSL.
//Retorno: Ponteiro de conexao com o banco de dados.
func ConectarComBancoDeDados(nomeUsuario, nomeBancoDeDados, senha, host, ssl string) *sql.DB {

	//Criando a string para a conecao com o banco de dados
	conexao := "user=" + nomeUsuario + "dbname=" + nomeBancoDeDados + "password=" + senha + "host=" + host + "sslmode=" + ssl

	//Abrindo a conecao com o banco de dados
	bd, erro := sql.Open("postgres", conexao)

	//Caso algum erro tenha ocorrido o programa ira abortar
	if erro != nil {
		panic(erro.Error)
	}

	//Retornando ponteiro de conexao com o banco de dados
	return bd
}
