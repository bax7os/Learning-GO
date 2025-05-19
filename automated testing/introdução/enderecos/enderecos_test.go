package enderecos

import "testing"

type cenarioDeTeste struct {
	enderecoInserido string
	tipoEsperado     string
}

// basic unit test

func TestTipoDeEndereco(t *testing.T) {
	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", "rua"},
		{"Avenida Paulista", "avenida"},
		{"Rodovia dos Imigrantes", "rodovia"},
		{"praça da Republica", "Inválido"},
	}

	for _, cenario := range cenariosDeTeste {
		tipoDeEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if tipoDeEnderecoRecebido != cenario.tipoEsperado {
			t.Errorf("Tipo de endereço esperado: %s, endereço recebido: %s", cenario.tipoEsperado, tipoDeEnderecoRecebido)
		}
	}

	//if tipoDeEnderecoRecebido != tipoDeEnderecoEsperado {
	//	t.Errorf("Tipo de endereço esperado: %s, endereço recebido: %s", tipoDeEnderecoEsperado, tipoDeEnderecoRecebido)
	//}

}
