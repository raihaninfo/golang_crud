package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/raihaninfo/golang_crud/views"
)

var (
	homeView   *views.View
	addView    *views.View
	ubdateView *views.View
	deleteView *views.View
)

func main() {
	homeView = views.NewView("views/fron-end/index.gohtml")
	addView = views.NewView("views/fron-end/add.gohtml")
	ubdateView = views.NewView("views/fron-end/update.gohtml")
	deleteView = views.NewView("views/fron-end/delete.gohtml")

	r := mux.NewRouter()
	r.PathPrefix("/asset/").Handler(http.StripPrefix("/asset/", http.FileServer(http.Dir("views/fron-end/asset"))))

	r.HandleFunc("/", home)
	r.HandleFunc("/add", add)
	r.HandleFunc("/update", update)
	r.HandleFunc("/delete", delete)

	fmt.Println("Listening port :8080")
	http.ListenAndServe(":8080", r)
}

func home(w http.ResponseWriter, r *http.Request) {
	err := homeView.Template.Execute(w, nil)
	FeatchError(err)
}

func add(w http.ResponseWriter, r *http.Request) {
	err := addView.Template.Execute(w, nil)
	FeatchError(err)
}

func update(w http.ResponseWriter, r *http.Request) {
	err := ubdateView.Template.Execute(w, nil)
	FeatchError(err)
}

func delete(w http.ResponseWriter, r *http.Request) {
	err := deleteView.Template.Execute(w, nil)
	FeatchError(err)
}

func FeatchError(err error) {
	if err != nil {
		panic(err)
	}
}
