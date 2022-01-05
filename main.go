package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Crud app")
	r := mux.NewRouter()
	r.PathPrefix("/asset/").Handler(http.StripPrefix("/asset/", http.FileServer(http.Dir("view/fron-end/asset"))))
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		temp, err := template.ParseFiles("view/fron-end/index.gohtml", "view/fron-end/header.gohtml")
		if err != nil {
			fmt.Println(err.Error())
		}
		temp.Execute(w, nil)
	})
	http.ListenAndServe(":8080", r)
}
