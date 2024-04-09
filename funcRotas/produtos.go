package funcRotas

import (
	"fmt"
	"github.com/Aleff5/golang-aulas/acaoProduto"
	"html/template"
	"net/http"
	"strconv"
)

// encapsula templates do site e devolve um template e um erro
var temp = template.Must(template.ParseGlob("src/golang-web/templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	listaProdutos := acaoProduto.BuscaProdutos()
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

		acaoProduto.CriaNvProduto(nome, descricao, precoConv, qtdConv) //CHAMA FUNC QUE CRIA PRODUTOS P ENVIAR P BD

	}
	http.Redirect(w, r, "/", 301)

}

func Deleta(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id")
	acaoProduto.DeletaProdutos(idDoProduto)
	http.Redirect(w, r, "/", 301)
}

func Editor(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id")
	produto := acaoProduto.Editor(idDoProduto)
	temp.ExecuteTemplate(w, "Edit", produto)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" { //BUSCA O QUE FOI DIGITADO (POSTADO) NA PAGINA.
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		idconv, err := strconv.Atoi(id)
		if err != nil {
			fmt.Println("erro na conversao do preço", err)
		}
		precoConv, err := strconv.ParseFloat(preco, 64) //REALIZA CONVERSAO DE TIPOS
		if err != nil {
			fmt.Println("erro na conversao do preço", err)
		}
		qtdConv, err := strconv.Atoi(quantidade) //REALIZA CONVERSAO DE TIPOS
		if err != nil {
			fmt.Println("erro na conversao da quantidade", err)
		}
		acaoProduto.Atualiza(idconv, nome, descricao, precoConv, qtdConv)
	}
	http.Redirect(w, r, "/", 301)
}
