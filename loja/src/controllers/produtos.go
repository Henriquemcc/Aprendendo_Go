package controllers

import (
	"html/template"
	"loja/models"
	"net/http"
)

//Criando o template da aplicacao web
var templateDaAplicacaoWeb = template.Must(template.ParseGlob("templates/*.html"))

//Esta funcao serve para atender a requisicao '/'
//Parametro: w: Instancia da interface usada pelo manipulador http para construir a resposta http.
//Parametro: r: Instancia da struct request, que representa uma requisicao recebida pelo servidor ou enviada pelo cliente.
func Index(w http.ResponseWriter, r *http.Request) {

	//Obtendo a lista de todos os produtos
	listaDeProdutos := models.BuscarTodosProdutos()

	//Executando aplicacao web
	templateDaAplicacaoWeb.ExecuteTemplate(w, "Index", listaDeProdutos)
}
