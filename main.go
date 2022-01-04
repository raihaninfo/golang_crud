package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Crud app")
	r := mux.NewRouter()
	http.ListenAndServe(":8080", r)
}
