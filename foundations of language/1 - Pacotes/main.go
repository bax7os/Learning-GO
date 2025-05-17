package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo da main")
	auxiliar.Escrever()
	erro := checkmail.ValidateFormat("1")
	fmt.Println(erro)
}
