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

	//Carregando rota para a pagina insert
	http.HandleFunc("/insert", controllers.Insert)

	//Carregando rota para o comando delete
	http.HandleFunc("/delete", controllers.Delete)

	//Carregando rota para a pagina editar
	http.HandleFunc("/edit", controllers.Edit)

	//Carregando rota para o comando atualizar
	http.HandleFunc("/update", controllers.Update)
}
