package main

import "fmt"

func diasDaSemana(numero int) string {
	var diaDaSemana string

	switch {
	case numero == 1:
		diaDaSemana = "Segunda-feira"
	case numero == 2:
		diaDaSemana = "Terça-feira"
	case numero == 3:
		diaDaSemana = "Quarta-feira"

	default:
		return "Número inválido"
	}
	return diaDaSemana
}

func main() {
	fmt.Println(diasDaSemana(1))
	fmt.Println(diasDaSemana(2))
	fmt.Println(diasDaSemana(3))
	fmt.Println(diasDaSemana(4))
	fmt.Println(diasDaSemana(5))
	fmt.Println(diasDaSemana(6))
	fmt.Println(diasDaSemana(7))
	fmt.Println(diasDaSemana(8))
}
