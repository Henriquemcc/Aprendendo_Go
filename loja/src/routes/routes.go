package routes

import (
	"loja/controllers"
	"net/http"
)

//CarregarRotas serve para carregar as rotas http do servidor web
func CarregarRotas() {
	http.HandleFunc("/", controllers.Index)
}
