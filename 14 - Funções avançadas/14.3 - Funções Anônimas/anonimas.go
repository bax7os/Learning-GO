package main

import "fmt"

func main() {
	func() {
		fmt.Println("Função Anônima")
	}()
	func(texto string) {
		fmt.Println(texto)
	}("oii")

}
