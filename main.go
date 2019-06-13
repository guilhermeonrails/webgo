package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float32
	Quantidade int
}

var tmpl = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	db := conectaComBancoDeDados()
	defer db.Close()
	http.HandleFunc("/", Index)
	http.ListenAndServe(":8000", nil)
}

func conectaComBancoDeDados() *sql.DB {
	constStr := "user=postgres dbname=alura_loja password=12345678 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", constStr)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func Index(w http.ResponseWriter, r *http.Request) {
	db := conectaComBancoDeDados()

	selectDB, err := db.Query("SELECT * FROM produtos")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	dados := []Produto{}

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
