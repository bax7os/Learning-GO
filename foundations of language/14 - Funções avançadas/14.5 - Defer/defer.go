package main

import "fmt"

func function1() {
	fmt.Println("Executando 1")
}
func function2() {
	fmt.Println("Executando 1")
}

func mediaAluno(nota1, nota2 float32) float32 {
	defer function1()
	defer function2()
	return (nota1 + nota2) / 2
}
