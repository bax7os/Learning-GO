package main

import (
	"fmt"
	"testes/enderecos"
)

func main() {
	tiposEnderecos := enderecos.TipoDeEndereco("cua 1")

	fmt.Println(tiposEnderecos)
}
