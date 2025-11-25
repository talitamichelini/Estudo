package main

import "fmt"

func main() {
	// BOOLEANO
	//fmt.Printf("Type: %T - Value: %v\n", true, true)

	// STRING
	//fmt.Printf("Type: %T - Value: %v\n", "Texto", "Texto")
	// texto ou número, estando entre aspas, é uma string

	// INTEIRO
	//fmt.Printf("Tipo: %T - Valor: %v\n", 100, 100)

	// DECIMAL
	fmt.Printf("Tipo: %T - Valor: %v\n", 10.95, 10.85)
}

// ANOTAÇÕES

//* As funções do pacote fmt sempre começam com letra maiúscula:
//fmt.Print
//fmt.Println
//fmt.Printf

// Prinf = imprimir formatação = placeholders
// os placeholders são os %T e %v, que são substituidos pelos valores passados depois da vírgula, os argumentos.
// %T = tipo - %v = valor (esses conseguimos encontrar no site do go na seção de "fmt")
// \n = nova linha; quebra de linha; enter

// o código passa 2 argumentos para ler:
// %T (placeholder) - para mostrar o tipo da variável (1º true ou texto ou etc)
// %v (placeholder) - para mostrar o valor da variável (2º true ou texto ou etc)

// Tipos principais:
// bool (booleano) true/false
// string (texto entre aspas)
// int (intero)
// float (decimal) [float32, float64]
