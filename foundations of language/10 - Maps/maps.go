package main

import "fmt"

func main() {
	usuario := map[string]string{
		"nome":      "João",
		"sobrenome": "Silva",
	}
	fmt.Println(usuario)

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "João",
			"ultimo":   "Silva",
		},
	}
	fmt.Println(usuario2)
}
