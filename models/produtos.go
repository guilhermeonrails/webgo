package models

import (
	"log"

	"github.com/guilhermelima/crudComGo/db"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodosProdutos() []Produto {
	db := db.ConectaComBancoDeDados()
	selectDB, err := db.Query("SELECT * FROM produtos")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectDB.Next() {
		var id int
		var nome, descricao string
		var preco float64
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

		produtos = append(produtos, p)

	}
	defer db.Close()
	return produtos
}

func CriaNovoProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaComBancoDeDados()

	insForm, err := db.Prepare("INSERT INTO produtos(nome, descricao, preco, quantidade) VALUES($1,$2,$3,$4)")
	if err != nil {
		panic(err.Error())
	}
	insForm.Exec(nome, descricao, preco, quantidade)
	log.Println("Inserindo novo produto:", nome, descricao, preco, quantidade)
	defer db.Close()
}

func DeletaProduto(id string) {
	db := db.ConectaComBancoDeDados()

	delForm, err := db.Prepare("delete from produtos WHERE id=$1")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(id)
	log.Println("Deletando produto com ID", id)
	defer db.Close()
}

func EditaProduto(id int) Produto {
	db := db.ConectaComBancoDeDados()

	produtoDoBanco, err := db.Query("SELECT * FROM produtos WHERE id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}

	for produtoDoBanco.Next() {
		var id int
		var nome, descricao string
		var preco float64
		var quantidade int

		err = produtoDoBanco.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade
	}
	defer db.Close()
	return p
}

func AtualizaProduto(id int, nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaComBancoDeDados()

	insForm, err := db.Prepare("update produtos set nome=$1, descricao=$2, preco=$3, quantidade=$4 WHERE id=$5")
	if err != nil {
		panic(err.Error())
	}
	insForm.Exec(nome, descricao, preco, quantidade, id)
	defer db.Close()
}
