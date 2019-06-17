package main

import (
	"net/http"

	"github.com/guilhermelima/crudComGo/controllers"
	_ "github.com/lib/pq"
)

func main() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
	http.ListenAndServe(":8000", nil)
}
