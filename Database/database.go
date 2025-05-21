package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	urlConexao := "golang:12345@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", urlConexao)
	if erro != nil {
		panic(erro)
	}

	fmt.Println(db)

	defer db.Close()
	if erro = db.Ping(); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println("Conexão aberta")

	linhas, erro := db.Query("select * from usuarios")

	fmt.Println(linhas)
	fmt.Println(erro)
	defer linhas.Close()

}
