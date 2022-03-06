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
	editView     *views.View
)

var store = sessions.NewCookieStore([]byte("secret-password"))

func init() {
	model.Dbcon()
}

var port string = ":8080"

func main() {
	homeView = views.NewView("views/front-end/index.gohtml")
	addView = views.NewView("views/front-end/add.gohtml")
	updateView = views.NewView("views/front-end/update.gohtml")
	deleteView = views.NewView("views/front-end/delete.gohtml")
	notFountView = views.NewView("views/front-end/notfount.gohtml")
	loginView = views.NewView("views/front-end/login.gohtml")
	editView = views.NewView("views/front-end/edit.gohtml")
	r := mux.NewRouter()
	r.PathPrefix("/asset/").Handler(http.StripPrefix("/asset/", http.FileServer(http.Dir("views/front-end/asset"))))
	r.NotFoundHandler = http.HandlerFunc(notFount)
	r.HandleFunc("/login", login)
	r.HandleFunc("/", home)
	r.HandleFunc("/add", add)
	r.HandleFunc("/addauth", addAuth)
	r.HandleFunc("/update", update)
	r.HandleFunc("/updateauth", updateAuth)
	r.HandleFunc("/update/{id}", updateId)
	r.HandleFunc("/delete", delete)
	r.HandleFunc("/delete/{id}", deleteId)

	fmt.Println("Listening port ", port)
	http.ListenAndServe(port, r)
}

type student struct {
	Id      string
	Name    string
	Address string
	Class   string
	Phone   string
}

var studentId string

func updateId(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	studentId = id
	stu := []student{}
	allStudent := model.ShowById(id)
	for i := 0; i < len(allStudent); i++ {
		id := allStudent[i]["id"].(string)
		name := allStudent[i]["name"].(string)
		address := allStudent[i]["address"].(string)
		class := allStudent[i]["class"].(string)
		phone := allStudent[i]["phone"].(string)
		stu = append(stu, student{Id: id, Name: name, Address: address, Class: class, Phone: phone})
	}

	err := editView.Template.Execute(w, stu)
	FetchError(err)
}

func deleteId(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	model.DeleteById(id)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func updateAuth(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	sname := r.FormValue("sname")
	saddress := r.FormValue("saddress")
	sclass := r.FormValue("sclass")
	sphone := r.FormValue("sphone")

	model.UpdateStudent(sname, saddress, sclass, sphone, studentId)
	http.Redirect(w, r, "/", http.StatusSeeOther)
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
	studen := []student{}
	allStudent := model.ShowAll()
	for i := 0; i < len(allStudent); i++ {
		id := allStudent[i]["id"].(string)
		name := allStudent[i]["name"].(string)
		address := allStudent[i]["address"].(string)
		class := allStudent[i]["class"].(string)
		phone := allStudent[i]["phone"].(string)
		studen = append(studen, student{Id: id, Name: name, Address: address, Class: class, Phone: phone})
	}
	err := homeView.Template.Execute(w, studen)
	FetchError(err)

}

func add(w http.ResponseWriter, r *http.Request) {
	err := addView.Template.Execute(w, nil)
	FetchError(err)

}

func addAuth(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.FormValue("sname")
	saddress := r.FormValue("saddress")
	class := r.FormValue("class")
	phone := r.FormValue("sphone")
	model.AddStudent(name, saddress, class, phone)
	http.Redirect(w, r, "/", http.StatusSeeOther)
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
