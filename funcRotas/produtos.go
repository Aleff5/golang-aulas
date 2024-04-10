package funcRotas

import (
	"encoding/json"
	"github.com/Aleff5/golang-aulas/acaoProduto"
	"html/template"
	"net/http"
)

// encapsula templates do site e devolve um template e um erro
var temp = template.Must(template.ParseGlob("src/golang-web/templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	listaProdutos := acaoProduto.BuscaProdutos()
	jsonData, err := json.Marshal(listaProdutos)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	_, err = w.Write(jsonData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	temp.ExecuteTemplate(w, "Index", listaProdutos)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	//if r.Method == "POST" { //BUSCA O QUE FOI DIGITADO (POSTADO) NA PAGINA.
	//	nome := r.FormValue("nome")
	//	descricao := r.FormValue("descricao")
	//	preco := r.FormValue("preco")
	//	quantidade := r.FormValue("quantidade")
	//
	//	precoConv, err := strconv.ParseFloat(preco, 64) //REALIZA CONVERSAO DE TIPOS
	//	if err != nil {
	//		fmt.Println("erro na conversao do preço", err)
	//	}
	//	qtdConv, err := strconv.Atoi(quantidade) //REALIZA CONVERSAO DE TIPOS
	//	if err != nil {
	//		fmt.Println("erro na conversao da quantidade", err)
	//	}
	//
	//	acaoProduto.CriaNvProduto(nome, descricao, precoConv, qtdConv) //CHAMA FUNC QUE CRIA PRODUTOS P ENVIAR P BD
	//
	//}
	if r.Method != http.MethodPost {
		http.Error(w, "Metodo invalido", http.StatusMethodNotAllowed)
	}

	var novoProduto acaoProduto.Produto
	if err := json.NewDecoder(r.Body).Decode(&novoProduto); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//inserindo
	acaoProduto.CriaNvProduto(novoProduto.Nome, novoProduto.Descricao, novoProduto.Preco, novoProduto.Quantidade)
	//sucesso
	w.WriteHeader(http.StatusCreated)
	//redireciona
	http.Redirect(w, r, "/", 301)

}

func Deleta(w http.ResponseWriter, r *http.Request) {

	//obtem ID
	idDoProduto := r.URL.Query().Get("id")

	//deletando
	acaoProduto.DeletaProdutos(idDoProduto)

	//sucesso
	w.WriteHeader(http.StatusOK)
	http.Redirect(w, r, "/", 301)
}

func Editor(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id")
	produto := acaoProduto.Editor(idDoProduto)

	//serializa dados p/ JSON
	jsonData, err := json.Marshal(produto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	temp.ExecuteTemplate(w, "Edit", produto)
}

func Update(w http.ResponseWriter, r *http.Request) {
	//if r.Method == "POST" { //BUSCA O QUE FOI DIGITADO (POSTADO) NA PAGINA.
	//	id := r.FormValue("id")
	//	nome := r.FormValue("nome")
	//	descricao := r.FormValue("descricao")
	//	preco := r.FormValue("preco")
	//	quantidade := r.FormValue("quantidade")
	//
	//	idconv, err := strconv.Atoi(id)
	//	if err != nil {
	//		fmt.Println("erro na conversao do preço", err)
	//	}
	//	precoConv, err := strconv.ParseFloat(preco, 64) //REALIZA CONVERSAO DE TIPOS
	//	if err != nil {
	//		fmt.Println("erro na conversao do preço", err)
	//	}
	//	qtdConv, err := strconv.Atoi(quantidade) //REALIZA CONVERSAO DE TIPOS
	//	if err != nil {
	//		fmt.Println("erro na conversao da quantidade", err)
	//	}
	//	acaoProduto.Atualiza(idconv, nome, descricao, precoConv, qtdConv)
	//}
	if r.Method != http.MethodPut {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	// Parse do corpo da requisição JSON para struct Produto
	var produtoAtualizado acaoProduto.Produto
	if err := json.NewDecoder(r.Body).Decode(&produtoAtualizado); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	acaoProduto.Atualiza(produtoAtualizado.Id, produtoAtualizado.Nome, produtoAtualizado.Descricao, produtoAtualizado.Preco, produtoAtualizado.Quantidade)

	//sucesso
	w.WriteHeader(http.StatusOK)
	http.Redirect(w, r, "/", 301)
}
