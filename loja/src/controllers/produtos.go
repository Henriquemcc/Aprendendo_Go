package controllers

import (
	"fmt"
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

	//Verificando se o metodo eh POST
	if r.Method == "POST" {

		//Obtendo valores do novo produto
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		//Convertendo o preco de string para float64
		precoConvertidoParaFloat64, erro := strconv.ParseFloat(preco, 64)

		//Caso algum erro ocorra, sera exibido uma mensagem de erro e a funcao sera abortada
		if erro != nil {
			log.Println("Erro na conversão do preço para float64:", erro)
			return
		}

		//Convertendo a quantidade de string para int
		quantidadeConvertidaParaInt, erro := strconv.Atoi(quantidade)

		//Caso algum erro ocorra, sera exibido uma mensagem de erro e a funcao sera abortada
		if erro != nil {
			log.Println("Erro na conversão da quantidade para int:", erro)
			return
		}

		//Criando instancia da classe Produto
		produto := models.Produto{}

		//Adicionando os valores lidos na instancia da classe Produto
		produto.SetNome(nome)
		produto.SetDescricao(descricao)
		sucesso, mensagemDeErro := produto.SetPreco(precoConvertidoParaFloat64)
		if !sucesso {
			fmt.Println(mensagemDeErro)
			return
		}
		sucesso, mensagemDeErro = produto.SetQuantidade(quantidadeConvertidaParaInt)
		if !sucesso {
			fmt.Println(mensagemDeErro)
			return
		}

		//Criando o novo produto
		models.CriarNovoProduto(produto)

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

	//Obtendo o id do produto
	idDoProduto := r.URL.Query().Get("id")

	//Obtendo os dados do produto
	produto := models.EditarProduto(idDoProduto)

	//Executndo aplicacao web
	templateDaAplicacaoWeb.ExecuteTemplate(w, "Edit", &produto)
}

//Update serve para atender a requisicao para o comando update, que ira atualizar o(s) valor(es) de um produto.
//Parametro: w: Instancia da interface usada pelo manipulador http para construir a resposta http.
//Parametro: r: Instancia da struct request, que representa uma requisicao recebida pelo servidor ou enviada pelo cliente.
func Update(w http.ResponseWriter, r *http.Request) {

	//Verificando se o metodo eh POST
	if r.Method == "POST" {

		//Obtendo valores do produto que foi editado
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		//Convertendo o id de string para int
		idConvertidoParaInt, erro := strconv.Atoi(id)

		//Caso algum erro ocorra, sera exibido uma mensagem de erro e a funcao sera abortada
		if erro != nil {
			log.Println("Erro na conversão do id para int:", erro)
			return
		}

		//Convertendo o preco de string para float64
		precoConvertidoParaFloat64, erro := strconv.ParseFloat(preco, 64)

		//Caso algum erro ocorra, sera exibido uma mensagem de erro e a funcao sera abortada
		if erro != nil {
			log.Println("Erro na conversão do preço para float64:", erro)
			return
		}

		//Convertendo a quantidade de string para int
		quantidadeConvertidaParaInt, erro := strconv.Atoi(quantidade)

		//Caso algum erro ocorra, sera exibido uma mensagem de erro e a funcao sera abortada
		if erro != nil {
			log.Println("Erro na conversão da quantidade para int:", erro)
			return
		}

		//Criando instancia da classe Produto
		produto := models.Produto{}

		//Adicionando os valores lidos na instancia da classe Produto
		sucesso, mensagemDeErro := produto.SetID(idConvertidoParaInt)
		if !sucesso {
			fmt.Println(mensagemDeErro)
			return
		}
		produto.SetNome(nome)
		produto.SetDescricao(descricao)
		sucesso, mensagemDeErro = produto.SetPreco(precoConvertidoParaFloat64)
		if !sucesso {
			fmt.Println(mensagemDeErro)
			return
		}
		sucesso, mensagemDeErro = produto.SetQuantidade(quantidadeConvertidaParaInt)
		if !sucesso {
			fmt.Println(mensagemDeErro)
			return
		}

		//Atualizando o produto
		models.AtualizarProduto(produto)
	}

	//Redirecionando usuario
	http.Redirect(w, r, "/", 301)

}
