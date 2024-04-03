package main

import (
	"github.com/Aleff5/golang-aulas/rotas"
	_ "github.com/lib/pq"
	"net/http"
)

func main() {
	rotas.CarregaRotas()
	//subir o servidor
	http.ListenAndServe(":8000", nil)
}
