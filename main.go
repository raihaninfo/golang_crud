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
	updateView   *views.View
	deleteView   *views.View
	notFountView *views.View
	loginView    *views.View
)

var store = sessions.NewCookieStore([]byte("secret-password"))

func init() {
	model.Dbcon()
}

var port string = ":8081"

func main() {
	homeView = views.NewView("views/front-end/index.gohtml")
	addView = views.NewView("views/front-end/add.gohtml")
	updateView = views.NewView("views/front-end/update.gohtml")
	deleteView = views.NewView("views/front-end/delete.gohtml")
	notFountView = views.NewView("views/front-end/notfount.gohtml")
	loginView = views.NewView("views/front-end/login.gohtml")
	r := mux.NewRouter()
	r.PathPrefix("/asset/").Handler(http.StripPrefix("/asset/", http.FileServer(http.Dir("views/front-end/asset"))))
	r.NotFoundHandler = http.HandlerFunc(notFount)
	r.HandleFunc("/login", login)
	r.HandleFunc("/", home)
	r.HandleFunc("/add", add)
	r.HandleFunc("/update", update)
	r.HandleFunc("/update/{id}", updateId)
	r.HandleFunc("/delete", delete)

	fmt.Println("Listening port ", port)
	http.ListenAndServe(port, r)
}

func updateId(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	res := struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	}{
		Id:   id,
		Name: "raihan",
	}

	// b, _ := json.Marshal(res)
	// w.Write(b)
	updateView.Template.Execute(w, res)
}

func login(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "login-session")
	FetchError(err)
	_, ok := session.Values["session"]
	if ok {
		http.Redirect(w, r, "/", http.StatusNotFound)
	}
	er := loginView.Template.Execute(w, nil)
	FetchError(er)
}

func home(w http.ResponseWriter, r *http.Request) {
	err := homeView.Template.Execute(w, nil)
	FetchError(err)
}

func add(w http.ResponseWriter, r *http.Request) {
	err := addView.Template.Execute(w, nil)
	FetchError(err)
}

func update(w http.ResponseWriter, r *http.Request) {
	err := updateView.Template.Execute(w, nil)
	FetchError(err)
}

func delete(w http.ResponseWriter, r *http.Request) {
	err := deleteView.Template.Execute(w, nil)
	FetchError(err)
}

func notFount(w http.ResponseWriter, r *http.Request) {
	err := notFountView.Template.Execute(w, nil)
	FetchError(err)
}

func FetchError(err error) {
	if err != nil {
		panic(err)
	}
}
