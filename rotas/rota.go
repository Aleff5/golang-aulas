package rotas

import (
	"github.com/Aleff5/golang-aulas/controllers"
	"net/http"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index) //função que é executada toda vez que se faz um request a raiz do servidor
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/insert", controllers.Insert)
	http.HandleFunc("/delete", controllers.Deleta)
	http.HandleFunc("/edit", controllers.Editor)
}
