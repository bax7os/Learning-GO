package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Idade int    `json:"idade"`
	Raca  string `json:"raca"`
}

func main() {
	cachorroEmJSON := `{"nome":"Lolla", "idade": 3, "raca":"Vira Lata"}`
	var c cachorro

	if erro := json.Unmarshal([]byte(cachorroEmJSON), &c); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(c)
}
