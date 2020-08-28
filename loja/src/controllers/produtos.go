package controllers

import (
	"html/template"
	"loja/models"
	"net/http"
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
