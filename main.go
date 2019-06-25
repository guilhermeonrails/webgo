package main

import (
	"net/http"

	"github.com/guilhermelima/crudComGo/routes"
	_ "github.com/lib/pq"
)

func main() {
	routes.CarregarRotas()
	http.ListenAndServe(":8000", nil)
}
