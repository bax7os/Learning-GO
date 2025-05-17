package main

import "fmt"

type pessoa struct {
	nome  string
	idade uint8
}
type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {

	p1 := pessoa{"João", 20}
	fmt.Println(p1)

	e1 := estudante{p1, "Engenharia", "USP"}
	fmt.Println(e1)

	fmt.Println("resultado")
}
