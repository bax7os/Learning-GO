package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func (p pessoa) salvar() {
	fmt.Printf("Salvando os dados de %s", p.nome)
}

func main() {
	pessoa1 := pessoa{"João", "Silva", 20}
	pessoa1.salvar()
}
