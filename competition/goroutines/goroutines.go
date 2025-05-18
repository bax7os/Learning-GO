package main

import (
	"fmt"
	"time"
)

func main() {

	// competition x parallelism
	// parallelism happens when we have multiple tasks running at the same time (more than one nucleus)
	// competition not necessary the execution at the same time

	go escrever("Olá mundo") //goroutine
	escrever("Programando em Go")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
