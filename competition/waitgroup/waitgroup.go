package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var waitGroup sync.WaitGroup
	waitGroup.Add(2) // two goroutines that he has to wait to finish

	go func() {

		escrever("Olá mundo")

		waitGroup.Done() // -1 of counter goroutines
	}()
	go func() {

		escrever("Programando em Go")
		waitGroup.Done()
	}()

	waitGroup.Wait()

}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
