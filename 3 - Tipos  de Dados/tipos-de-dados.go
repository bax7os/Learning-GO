package main

import (
	"errors"
	"fmt"
)

func main() {
	var num int = 1000
	fmt.Println(num)

	var num2 uint = 1000
	fmt.Println(num2)

	// alias
	// INT32 = RUNE
	var num3 rune = 1234
	fmt.Println(num3)

	// alias
	// UINT8 = BYTE
	var num4 byte = 123
	fmt.Println(num4)

	var erro error = errors.New("Descrição do erro")
	fmt.Println(erro)
}
