package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	_ "github.com/mattn/go-sqlite3"
	"github.com/raihaninfo/golang_crud/model"
	"github.com/raihaninfo/golang_crud/views"
)

var (
	homeView     *views.View
	addView      *views.View
	ubdateView   *views.View
	deleteView   *views.View
	notFountView *views.View
	loginView    *views.View
)

var store = sessions.NewCookieStore([]byte("secret-password"))

func init() {
	model.Dbcon()
}

func main() {
	homeView = views.NewView("views/fron-end/index.gohtml")
	addView = views.NewView("views/fron-end/add.gohtml")
	ubdateView = views.NewView("views/fron-end/update.gohtml")
	deleteView = views.NewView("views/fron-end/delete.gohtml")
	notFountView = views.NewView("views/fron-end/notfount.gohtml")
	loginView = views.NewView("views/fron-end/login.gohtml")
	r := mux.NewRouter()
	r.PathPrefix("/asset/").Handler(http.StripPrefix("/asset/", http.FileServer(http.Dir("views/fron-end/asset"))))
	r.NotFoundHandler = http.HandlerFunc(notFount)
	r.HandleFunc("/login", login)
	r.HandleFunc("/", home)
	r.HandleFunc("/add", add)
	r.HandleFunc("/update", update)
	r.HandleFunc("/delete", delete)

	fmt.Println("Listening port :8080")
	http.ListenAndServe(":8080", r)
}

func login(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "login-session")
	FeatchError(err)
	_, ok := session.Values["session"]
	if ok {
		http.Redirect(w, r, "/", http.StatusNotFound)
	}
	er := loginView.Template.Execute(w, nil)
	FeatchError(er)
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

func notFount(w http.ResponseWriter, r *http.Request) {
	err := notFountView.Template.Execute(w, nil)
	FeatchError(err)
}

func FeatchError(err error) {
	if err != nil {
		panic(err)
	}
}
