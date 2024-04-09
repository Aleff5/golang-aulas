package acaoProduto

import (
	"github.com/Aleff5/golang-aulas/db"
)

type Produto struct {
	Id         int     `json:"id"`
	Nome       string  `json:"nome"`
	Descricao  string  `json:"descricao"`
	Preco      float64 `json:"preco"`
	Quantidade int     `json:"quantidade"`
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

// FUNC QUE CRIA PRODUTOS BASEADO NA ENTRADADA SITE
// recebe dados do front
func CriaNvProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaDB()

	enviaDados, err := db.Prepare("insert into produtos(nome, descricao, preco, quantidade) values ($1,$2,$3,$4)") //escreve cod sql

	if err != nil {
		panic(err.Error())
	}

	enviaDados.Exec(nome, descricao, preco, quantidade)
	defer db.Close()
}

// FUNC QUE DELETA BASEADO NA ENTRADA DO SITE
// recebe dados do front
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
		produtoAtualizar.Id = id
		produtoAtualizar.Nome = nome
		produtoAtualizar.Descricao = descricao
		produtoAtualizar.Preco = preco
		produtoAtualizar.Quantidade = quantidade
	}
	defer db.Close()
	return produtoAtualizar
}

func Atualiza(id int, nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaDB()

	atualizador, err := db.Prepare("update produtos set nome=$1, descricao=$2, preco=$3, quantidade=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}
	atualizador.Exec(nome, descricao, preco, quantidade, id)
	defer db.Close()
}
