package routes

import (
	"loja/controllers"
	"net/http"
)

//CarregarRotas serve para carregar e direcionar as rotas http do servidor web
func CarregarRotas() {

	//Carregando rota para a pagina index
	http.HandleFunc("/", controllers.Index)

	//Carregando rota para a pagina new
	http.HandleFunc("/new", controllers.New)

	//Carregando rote para a pagina insert
	http.HandleFunc("/insert", controllers.Insert)
}
