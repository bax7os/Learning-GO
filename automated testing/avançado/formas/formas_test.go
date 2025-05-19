package formas

import "testing"

func TestAormas(t *testing.T) {
	t.Run("Retangulo", func(t *testing.T) {
		ret := Retangulo{10, 12}
		areaEsperada := float64(120)
		areaRecebida := ret.Area()

		if areaEsperada != areaRecebida {
			t.Fatalf("Esperava %f mas recebeu %f", areaEsperada, areaRecebida)
		}
	})
	t.Run("Circulo", func(t *testing.T) {
		circ := Circulo{10}
		areaEsperada := float64(314)
		areaRecebida := circ.Area()

		if areaEsperada != areaRecebida {
			t.Errorf("Esperava %f mas recebeu %f", areaEsperada, areaRecebida)
		}
	})

}
