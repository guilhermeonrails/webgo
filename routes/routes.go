package routes

import (
	"net/http"

	"github.com/guilhermelima/crudComGo/controllers"
)

func CarregarRotas() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/insert", controllers.Insert)
	http.HandleFunc("/delete", controllers.Delete)
}
