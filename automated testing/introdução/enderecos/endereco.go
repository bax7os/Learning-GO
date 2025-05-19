package enderecos

import (
	"fmt"
	"strings"
)

func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}

	enderecoEmLetraMinuscula := strings.ToLower(endereco)
	primeiraPalavra := strings.Split(enderecoEmLetraMinuscula, " ")[0]

	valido := false
	fmt.Println(primeiraPalavra)
	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavra {
			valido = true
			break
		}
	}

	if valido {
		return primeiraPalavra
	}
	return "Inválido"

}
