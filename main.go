package main

import (
	"net/http"

	"github.com/guilhermelima/crudComGo/routes"
	_ "github.com/lib/pq"
)

func main() {
	routes.CarregarRotas()
	http.ListenAndServe(":3001", nil)
}
