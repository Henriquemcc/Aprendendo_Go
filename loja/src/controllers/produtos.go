package controllers

import (
	"html/template"
	"log"
	"loja/models"
	"net/http"
	"strconv"
)

//Criando o template da aplicacao web
var templateDaAplicacaoWeb = template.Must(template.ParseGlob("templates/*.html"))

//Index serve para atender a requisicao para a pagina Index, que mostra os produtos da loja.
//Parametro: w: Instancia da interface usada pelo manipulador http para construir a resposta http.
//Parametro: r: Instancia da struct request, que representa uma requisicao recebida pelo servidor ou enviada pelo cliente.
func Index(w http.ResponseWriter, r *http.Request) {

	//Obtendo a lista de todos os produtos
	listaDeProdutos := models.BuscarTodosProdutos()

	//Executando aplicacao web para a pagina Index
	templateDaAplicacaoWeb.ExecuteTemplate(w, "Index", listaDeProdutos)
}

//New serve para atender a requisicao para a pagina New, que permite adicionar produtos na loja.
//Parametro: w: Instancia da interface usada pelo manipulador http para construir a resposta http.
//Parametro: r: Instancia da struct request, que representa uma requisicao recebida pelo servidor ou enviada pelo cliente.
func New(w http.ResponseWriter, r *http.Request) {

	//Executand aplicacao web para a pagina New
	templateDaAplicacaoWeb.ExecuteTemplate(w, "New", nil)
}

//Insert serve para atender a requisicao para a pagina Insert, que ira inserir um produto no banco de dados da loja.
//Parametro: w: Instancia da interface usada pelo manipulador http para construir a resposta http.
//Parametro: r: Instancia da struct request, que representa uma requisicao recebida pelo servidor ou enviada pelo cliente.
func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		//Obtendo valores do novo produto
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		//Convertendo o preco de string para float64
		precoConvertidoParaFloat64, erro := strconv.ParseFloat(preco, 64)

		//Caso algum erro ocorra, sera exibido uma mensagem de erro
		if erro != nil {
			log.Println("Erro na conversão do preço:", erro)
		}

		//Convertendo a quantidade de string para int
		quantidadeConvertidaParaInt, erro := strconv.Atoi(quantidade)

		//Caso algum erro ocorra, sera exibido uma mensagem de erro
		if erro != nil {
			log.Println("Erro na conversão da quantidade:", erro)
		}

		//Criando novo produto
		models.CriarNovoProduto(nome, descricao, precoConvertidoParaFloat64, quantidadeConvertidaParaInt)

	}

	//Redirecionando usuario
	http.Redirect(w, r, "/", 301)
}

//Delete serve para atender a requisicao para o comando Delete, que ira deletar um produto no banco de dados da loja.
//Parametro: w: Instancia da interface usada pelo manipulador http para construir a resposta http.
//Parametro: r: Instancia da struct request, que representa uma requisicao recebida pelo servidor ou enviada pelo cliente.
func Delete(w http.ResponseWriter, r *http.Request) {

	//Obtendo o id do produto
	idDoProduto := r.URL.Query().Get("id")

	//Deletando o produto
	models.DeletarProduto(idDoProduto)

	//Redirecionando usuario
	http.Redirect(w, r, "/", 301)
}

//Edit serve para atender a requisicao para a pagina Edit, que ira realizar a edicao de um produto.
//Parametro: w: Instancia da interface usada pelo manipulador http para construir a resposta http.
//Parametro: r: Instancia da struct request, que representa uma requisicao recebida pelo servidor ou enviada pelo cliente.
func Edit(w http.ResponseWriter, r *http.Request) {
	templateDaAplicacaoWeb.ExecuteTemplate(w, "Edit", nil)
}
