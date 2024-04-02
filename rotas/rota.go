package rotas

import (
	"../controllers"
	"net/http"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index) //função que é executada toda vez que se faz um request a raiz do servidor
}
