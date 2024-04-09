package rotas

import (
	"github.com/Aleff5/golang-aulas/funcRotas"
	"net/http"
)

func CarregaRotas() {
	http.HandleFunc("/", funcRotas.Index) //função que é executada toda vez que se faz um request a raiz do servidor
	http.HandleFunc("/new", funcRotas.New)
	http.HandleFunc("/insert", funcRotas.Insert)
	http.HandleFunc("/delete", funcRotas.Deleta)
	http.HandleFunc("/edit", funcRotas.Editor)
	http.HandleFunc("/update", funcRotas.Update)
}
