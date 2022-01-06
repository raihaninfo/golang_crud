package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

var (
	homeTemplate *template.Template
)

func main() {
	var err error
	homeTemplate, err = template.ParseFiles("view/view.gohtml")
	if err != nil {
		panic(err)
	}
	fmt.Println("Listening port :8080")
	r := mux.NewRouter()
	r.PathPrefix("/asset/").Handler(http.StripPrefix("/asset/", http.FileServer(http.Dir("view/fron-end/asset"))))
	r.HandleFunc("/", home)
	http.ListenAndServe(":8080", r)
}

func home(w http.ResponseWriter, r *http.Request) {
	err := homeTemplate.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}
