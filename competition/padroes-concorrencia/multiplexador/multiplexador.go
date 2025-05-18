package main

import (
	"fmt"
	"time"
)

func main() {

	canal := multiplexar(escrever("Ola Mundo"), escrever("Programando em Go"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}

}
func multiplexar(canal1, canal2 <-chan string) <-chan string {
	canal := make(chan string)
	go func() {
		for {
			select {
			case mensagem := <-canal1:
				canal <- mensagem
			case mensagem := <-canal2:
				canal <- mensagem
			}
		}
	}()
	return canal
}
func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("%s", texto)
			time.Sleep(time.Millisecond * 500)
		}
	}()
	return canal
}
