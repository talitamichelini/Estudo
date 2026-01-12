package main

import "fmt"

// condicionais if-else em Go
// IF - Se
// Else - Senão

func main() {

	//valor := 3 // esse é o valor base
	// // if (condição) { bloco de código }
	// if valor == 1 { // se, essa é a condição
	// 	fmt.Println("O valor é igual a 1")
	// } else { // se não, está "dentro" do if
	// 	fmt.Println("O valor é diferente de 1")
	// }

	// if valor == 1 {
	// 	fmt.Println("O valor é igual a 1")
	// } else if valor == 2 { // senão se, outra condição
	// 	fmt.Println("O valor é igual a 2")
	// } else { // se não for nenhum dos dois
	// 	fmt.Println("O valor é diferente de 1 e 2")
	// }

	// aplicando Operações
	// fmt.Println(1 + 1) // soma
	// fmt.Println(2 - 1) // subtração
	// fmt.Println(2 * 2) // multiplicação
	// fmt.Println(4 / 2) // divisão
	// fmt.Println(5 % 2) // módulo (resto da divisão)

	// para descobrir se o número é par ou ímpar
	// é só saber se é divisível por 2
	numero := 7
	if numero%2 == 0 { // = 7%2, porcento = resto da divisão
		// significa que o resto é 0, ímpar
		fmt.Println("é par") // pode usar atalho com %d
	} else {
		fmt.Println("é ímpar") // pode usar atalho com %d
	}

}
