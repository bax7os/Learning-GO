package auxiliar

import "fmt"

// If a function starts with a capital letter, it can be exported
// If a function not starts with a capital letter, it can't be exported
// This happens because the golang is not a language oriented to OOP
func Escrever() {
	fmt.Println("Escrevendo do pacote auxiliar")
}
