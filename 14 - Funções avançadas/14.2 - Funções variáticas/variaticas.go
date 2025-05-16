package main

import "fmt"

func soma(numeros ...int) int {
	total := 0
	for _, num := range numeros {
		total += num
	}
	return total
}

func escrever(texto string, numeros ...int) {
	for _, num := range numeros {
		fmt.Println(texto, num)
	}
}

func main() {
	totalSoma := soma(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(totalSoma)
	escrever("Escrevendo", 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
}
