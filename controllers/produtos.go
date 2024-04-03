package controllers

import (
	"fmt"
	"github.com/Aleff5/golang-aulas/modelos"
	"html/template"
	"net/http"
	"strconv"
)

// encapsula templates do site e devolve um template e um erro
var temp = template.Must(template.ParseGlob("src/golang-web/templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	listaProdutos := modelos.BuscaProdutos()
	temp.ExecuteTemplate(w, "Index", listaProdutos)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" { //BUSCA O QUE FOI DIGITADO (POSTADO) NA PAGINA.
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConv, err := strconv.ParseFloat(preco, 64) //REALIZA CONVERSAO DE TIPOS
		if err != nil {
			fmt.Println("erro na conversao do preço", err)
		}
		qtdConv, err := strconv.Atoi(quantidade) //REALIZA CONVERSAO DE TIPOS
		if err != nil {
			fmt.Println("erro na conversao da quantidade", err)
		}

		modelos.CriaNvProduto(nome, descricao, precoConv, qtdConv) //CHAMA FUNC QUE CRIA PRODUTOS P ENVIAR P BD

	}
	http.Redirect(w, r, "/", 301)

}

func Deleta(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id")
	modelos.DeletaProdutos(idDoProduto)
	http.Redirect(w, r, "/", 301)
}

func Editor(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id")
	produto := modelos.Editor(idDoProduto)
	temp.ExecuteTemplate(w, "Edit", produto)
}
