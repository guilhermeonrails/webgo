package controllers

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/guilhermelima/crudComGo/models"
)

var tmpl = template.Must(template.ParseGlob("templates/*"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := models.BuscaTodosProdutos()
	tmpl.ExecuteTemplate(w, "Index", todosOsProdutos)
}

func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoF, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversão:", err)
		}

		quantidadeI, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversão:", err)
		}

		models.CriaNovoProduto(nome, descricao, precoF, quantidadeI)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id")
	models.DeletaProduto(idDoProduto)
	http.Redirect(w, r, "/", 301)
}
