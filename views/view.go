package views

import (
	"log"
	"text/template"
)

type View struct {
	Template *template.Template
}

func NewView(files ...string) *View {
	files = append(files, "views/front-end/header.gohtml", "views/front-end/menu.gohtml")
	t, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err)
	}
	return &View{
		Template: t,
	}
}
