package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Home"))
	})
	http.HandleFunc("/mensagem", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Mensagem"))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
