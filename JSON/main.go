package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Idade int    `json:"idade"`
	Raca  string `json:"raca"`
}

func main() {
	c := cachorro{"Lolla", 3, "Vira Lata"}
	fmt.Println(c)

	cachorroEmJSON, erro := json.Marshal(c)
	if erro != nil {
		fmt.Println(erro)
	}
	fmt.Println(string(cachorroEmJSON))

	bytes.NewBuffer(cachorroEmJSON)
}
