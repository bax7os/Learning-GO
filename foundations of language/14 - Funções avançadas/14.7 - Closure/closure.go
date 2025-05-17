package main

import "fmt"

func closure() func() {
	texto := "Dentro da closure"
	funcao := func() {
		fmt.Println(texto)
	}

	return funcao
}

func main() {
	texto := "Dentro da main"
	fmt.Println(texto)
	novaFuncao := closure()
	novaFuncao()
}
