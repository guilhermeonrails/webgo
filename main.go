package main

import (
	"html/template"
	"net/http"
)

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float32
	Quantidade int
}

var tmpl = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", Index)
	http.ListenAndServe(":8000", nil)
}

func Index(w http.ResponseWriter, r *http.Request) {
	dados := []Produto{
		{Nome: "Camiseta", Descricao: "Muito bonita", Preco: 39.9, Quantidade: 10},
		{Nome: "Sapato", Descricao: "Para ocasiões formais", Preco: 199.99, Quantidade: 5},
		{Nome: "Notebook", Descricao: "Muito rápido", Preco: 2099.99, Quantidade: 1},
		{Nome: "Fone", Descricao: "Confortável e boa sonoridade", Preco: 19.99, Quantidade: 10},
	}
	tmpl.ExecuteTemplate(w, "Index", dados)
}
