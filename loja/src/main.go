package main

import (
	"loja/produtos"
	"net/http"
	"text/template"
)

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

	//Criando uma lista de produtos
	var listaDeProdutos []produtos.Produto

	//Criando o primeiro produto e adicionando a lista
	camiseta := produtos.Produto{}
	camiseta.SetNome("camiseta")
	camiseta.SetDescricao("Azul")
	camiseta.SetPreco(20.00)
	camiseta.SetQuantidade(5)
	listaDeProdutos = append(listaDeProdutos, camiseta)

	//Criando o segundo produto e adicionando a lista
	computadorGamer := produtos.Produto{}
	computadorGamer.SetNome("Computador Gamer")
	computadorGamer.SetDescricao("Computador Gamer poderoso")
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
