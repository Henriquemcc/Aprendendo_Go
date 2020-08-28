package main

import (
	_ "loja/models"
	"loja/routes"
	"net/http"
)

//Esta funcao eh a funcao principal, onde o programa comeca a ser executado.
func main() {

	routes.CarregarRotas()
	http.ListenAndServe(":8000", nil)
}
