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

//Insert serve para atender a requisicao para a pagina Insert, que ira inserir os produtos no banco de dados da loja.
//Parametro: w: Instancia da interface usada pelo manipulador http para construir a resposta http.
//Parametro: r: Instancia da struct request, que representa uma requisicao recebida pelo servidor ou enviada pelo cliente.
func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertidoParaFloat64, erro := strconv.ParseFloat(preco, 64)

		if erro != nil {
			log.Println("Erro na conversão do preço:", erro)
		}

		quantidadeConvertidaParaInt, erro := strconv.Atoi(quantidade)

		if erro != nil {
			log.Println("Erro na conversão da quantidade:", erro)
		}

		models.CriarNovoProduto(nome, descricao, precoConvertidoParaFloat64, quantidadeConvertidaParaInt)

	}

	http.Redirect(w, r, "/", 301)
}
