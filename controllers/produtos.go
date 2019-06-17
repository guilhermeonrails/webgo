package controllers

import (
	"net/http"
	"text/template"

	"github.com/guilhermelima/crudComGo/db"
	"github.com/guilhermelima/crudComGo/models"
)

var tmpl = template.Must(template.ParseGlob("templates/*"))

func Index(w http.ResponseWriter, r *http.Request) {
	db := db.ConectaComBancoDeDados()

	selectDB, err := db.Query("SELECT * FROM produtos")
	if err != nil {
		panic(err.Error())
	}

	p := models.Produto{}
	dados := []models.Produto{}

	for selectDB.Next() {
		var id int
		var nome, descricao string
		var preco float32
		var quantidade int

		err = selectDB.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		// Junta a Struct com Array
		dados = append(dados, p)
	}
	tmpl.ExecuteTemplate(w, "Index", dados)
	defer db.Close()
}

func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}
