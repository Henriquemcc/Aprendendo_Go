package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
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
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

//Esta funcao serve para atender a requisicao '/'
//Parametro: w: Instancia da interface usada pelo manipulador http para construir a resposta http.
//Parametro: r: Instancia da struct request, que representa uma requisicao recebida pelo servidor ou enviada pelo cliente.
func index(w http.ResponseWriter, r *http.Request) {

	//Conectando com o banco de dados
	bd := ConectarComBancoDeDados(ObterCredenciaisDeAcessoAoBancoDeDados(nomeArquivoCredencialBancoDeDados))

	//Obtendo todos os produtos do banco de dados
	selecaoDeTodosProdutos, erro := bd.Query("select * from produtos")

	//Abortando o programa caso algum erro ocorra
	if erro != nil {
		fmt.Println("Erro ao conectar com o banco de dados:", erro)
		os.Exit(-1)
	}

	listaDeProdutos := []produtos.Produto{}

	for selecaoDeTodosProdutos.Next() {

		//Criando as variaveis que irao receber os valores do banco de dados
		var id, quantidade int
		var nome, descricao string
		var preco float64

		//Lendo os dados do banco de dados
		erro = selecaoDeTodosProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)

		//Caso algum erro acontecer o programa ira abortar
		if erro != nil {
			fmt.Println("Erro ao obter os dados do banco de dados:", erro)
			os.Exit(-1)
		}

		//Criando uma instancia de produto
		produto := produtos.Produto{}

		//Adicionando os dados lidos do banco de dados a instancia da struct Produto
		produto.SetID(id)
		produto.SetNome(nome)
		produto.SetDescricao(descricao)
		produto.SetPreco(preco)
		produto.SetQuantidade(quantidade)

		//Adicionando a instancia da struct Produto a slice de produtos
		listaDeProdutos = append(listaDeProdutos, produto)

	}

	//Executando aplicacao web
	templateDaAplicacaoWeb.ExecuteTemplate(w, "Index", listaDeProdutos)

	//Fechando a conexao com o banco de dados
	erro = bd.Close()

	if erro != nil {
		fmt.Println("Um erro ocorreu ao tentar fechar a conexao com o banco de dados:", erro)
		os.Exit(-1)
	}
}

//ObterCredenciaisDeAcessoAoBancoDeDados serve para obter as credenciais de acesso ao banco de dados de um arquivo.
//Parametro: nomeArquivo: Nome do arquivo em que sera lido para obter as credenciais de acesso ao banco de dados.
//Retorno: string: Nome de usuario do banco de dados.
//Retorno: string: Nome do banco de dados.
//Retorno: string: Senha do banco de dados
//Retorno: string: Host do banco de dados.
//Retorno: string: Modo de configuracao do protocolo SSL do banco de dados.
func ObterCredenciaisDeAcessoAoBancoDeDados(nomeArquivo string) (string, string, string, string, string) {

	//Lendo todos os dados do arquivo
	d, erro := ioutil.ReadFile(nomeArquivo)

	//Abortando a execucao do programa caso algum erro ocorra
	if erro != nil {
		fmt.Println("Um erro ocorreu ao ler os dados do arquivo de credenciais:", erro)
		os.Exit(-1)
	}

	//Convertendo os bytes em string e separando os dados lidos
	dadosArquivo := strings.Split(string(d), "\n")

	//Obtendo o nome de usuario
	nomeUsuario := dadosArquivo[0]

	//Obtendo o nome do banco de dados
	nomeBancoDeDados := dadosArquivo[1]

	//Obtendo a senha do banco de dados
	senha := dadosArquivo[2]

	//Obtendo o host do banco de dados
	host := dadosArquivo[3]

	//Obtendo o modo de ssl do banco de dados
	ssl := dadosArquivo[4]

	//Retornando as credenciais
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
	conexao := "user=" + nomeUsuario + " dbname=" + nomeBancoDeDados + " password=" + senha + " host=" + host + " sslmode=" + ssl

	//Abrindo a conecao com o banco de dados
	bd, erro := sql.Open("postgres", conexao)

	//Caso algum erro tenha ocorrido o programa ira abortar
	if erro != nil {
		fmt.Println("Um erro ocorreu ao conectar com o banco de dados:", erro)
		os.Exit(-1)
	}

	//Retornando ponteiro de conexao com o banco de dados
	return bd
}
