package formas

type Retangulo struct {
	Largura float64
	Altura  float64
}

func (r Retangulo) Area() float64 {
	return r.Largura * r.Altura
}

type Circulo struct {
	radius float64
}

func (c Circulo) Area() float64 {
	return 3.14 * c.radius * c.radius
}

type Forma interface {
	Area() float64
}
