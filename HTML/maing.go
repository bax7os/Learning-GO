package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type usuario struct {
	Nome      string
	Sobrenome string
}

var templates *template.Template

func main() {

	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		u := usuario{"João", "Silva"}
		templates.ExecuteTemplate(w, "home.html", u)
	})
	fmt.Println("Escutando na porta 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
