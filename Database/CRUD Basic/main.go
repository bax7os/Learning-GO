package main

import (
	"crud/servidor"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/usuarios", servidor.CriaUsuario).Methods("POST")
	router.HandleFunc("/usuarios", servidor.BuscarUsuarios).Methods("GET")
	router.HandleFunc("/usuarios/{id}", servidor.BuscarUsuario).Methods("GET")
	router.HandleFunc("/usuarios/{id}", servidor.AtualizarUsuario).Methods("PUT")
	router.HandleFunc("/usuarios/{id}", servidor.DeletarUsuario).Methods("DELETE")
	fmt.Println("Escutando na porta 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
