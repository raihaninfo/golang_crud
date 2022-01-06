package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/raihaninfo/golang_crud/views"
)

var (
	homeView *views.View
)

func main() {
	homeView = views.NewView("views/fron-end/index.gohtml")

	fmt.Println("Listening port :8080")
	r := mux.NewRouter()
	r.PathPrefix("/asset/").Handler(http.StripPrefix("/asset/", http.FileServer(http.Dir("views/fron-end/asset"))))
	r.HandleFunc("/", home)
	http.ListenAndServe(":8080", r)
}

func home(w http.ResponseWriter, r *http.Request) {
	err := homeView.Template.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}
