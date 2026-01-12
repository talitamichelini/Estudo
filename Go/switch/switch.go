package main

import "fmt"

func main() {

	// outra forma de condicional

	posicao := 2
	switch posicao {
	case 1: // da as opções de casos
		fmt.Println("Primeiro lugar")
	case 2:
		fmt.Println("Segundo lugar")
	case 3:
		fmt.Println("Terceiro lugar")
	default: // padrão caso nenhum dos outros seja satisfeito
		fmt.Println("Não está no pódio")

	}
	// pode ser aplicado com os outros tipos também, assim como if/else if

	nome := "Bob"
	switch nome {
	case "Talita":
		fmt.Println("Olá Talita")
	case "Bob":
		fmt.Println("Olá Bob")
	default:
		fmt.Println("Olá estranho")
	}
}
