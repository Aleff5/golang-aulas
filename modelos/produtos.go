package modelos

import (
	"github.com/Aleff5/golang-aulas/db"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

// FUNC QUE BUSCA PRODUTOS NO BD PARA ENVIA-LOS P SITE
func BuscaProdutos() []Produto {
	db := db.ConectaDB()

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

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}
	defer db.Close()
	return produtos

}

// FUNC QUE CRIA PRODUTOS BASEADO NA ENTRADADO SITE
func CriaNvProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaDB()

	enviaDados, err := db.Prepare("insert into produtos(nome, descricao, preco, quantidade) values ($1,$2,$3,$4)") //escreve cod sql

	if err != nil {
		panic(err.Error())
	}

	enviaDados.Exec(nome, descricao, preco, quantidade)
	defer db.Close()
}

func DeletaProdutos(id string) {
	db := db.ConectaDB()

	deletador, err := db.Prepare("delete from produtos where id=$1")
	if err != nil {
		panic(err.Error())
	}
	deletador.Exec(id)
	defer db.Close()
}

func Editor(id string) Produto {
	db := db.ConectaDB()
	produtoDb, err := db.Query("select * from produtos where id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	produtoAtualizar := Produto{}
	for produtoDb.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = produtoDb.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		produtoAtualizar.Nome = nome
		produtoAtualizar.Descricao = descricao
		produtoAtualizar.Preco = preco
		produtoAtualizar.Quantidade = quantidade
	}
	defer db.Close()
	return produtoAtualizar
}
