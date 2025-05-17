package main

import "fmt"

type retangulo struct {
	largura float64
	altura  float64
}

func (r retangulo) area() float64 {
	return r.largura * r.altura
}

type circulo struct {
	radius float64
}

func (c circulo) area() float64 {
	return 3.14 * c.radius * c.radius
}

type forma interface {
	area() float64
}

func escreverArea(f forma) {
	fmt.Println(f.area())
}

func main() {
	retangulo := retangulo{10, 15}

	escreverArea(retangulo)

	c := circulo{10}
	escreverArea(c)
}
