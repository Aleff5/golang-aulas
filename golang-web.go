package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"html/template"
	"net/http"
)

func conectaDB() *sql.DB {
	conexao := "user=postgres dbname=loja password=comida host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

// encapsula templates do site e devolve um template e um erro
var temp = template.Must(template.ParseGlob("src/golang-web/templates/*.html"))

func main() {

	//função que é executada toda vez que se faz um request a raiz do servidor
	http.HandleFunc("/", index)
	//subir o servidor
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	db := conectaDB()

	selecaoProdutos, err := db.Query("select * from produtos")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selecaoProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selecaoProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}

	temp.ExecuteTemplate(w, "Index", produtos)
	defer db.Close()
}
