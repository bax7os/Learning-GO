package main

import "fmt"

type usuario struct {
	nome      string
	sobrenome string
	idade     uint8
	ativo     bool
	endereco  endereco
}
type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	var u usuario
	u.nome = "João"
	u.idade = 20
	u.ativo = true
	u.sobrenome = "Silva"

	fmt.Println(u)

	u2 := usuario{"Ana", "Bastos", 20, true, endereco{"Rua 1", 10}}
	fmt.Println(u2)
}
