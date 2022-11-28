package render

import (
	"log"
	"net/http"
	"text/template"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	template, _ := template.ParseFiles("./templates/" + tmpl)
	err := template.Execute(w, nil)

	if err != nil {
		log.Fatal(err)
		return
	}

}
