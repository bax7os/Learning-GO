package main

import (
	"fmt"
)

func main() {
	i := 0
	for i < 10 {
		i++
		fmt.Println("Incrementando", i)
		//time.Sleep(time.Second)

	}

	for j := 0; j < 10; j++ {
		fmt.Println("Incrementando", j)
		//time.Sleep(time.Second)
	}

	nomes := [3]string{"João", "Davi", "Maria"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}
}
