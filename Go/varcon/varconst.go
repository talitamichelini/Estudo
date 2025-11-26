package main

import "fmt"

// tem que iniciar na linha 1

func main() {
	//VARIÁVEIS
	// var nome string
	// nome = "Maria"
	// fmt.Println(nome)

	//var a, b int
	// pode separar os argumentos por vírgula
	// a = 3
	// b = 7
	// fmt.Println(a + b)
	// me lembra problemas de 2º grau
	//ou
	// var a, b int = 3, 7
	// fmt.Println(a)
	// fmt.Println(b)

	// var d = true
	// fmt.Println(d)

	// OU - SIMPLIFICANDO
	// ao invés de usar var, pode usar :=
	// número := 1
	// fmt.Println(número)
	// nome := "João"
	// fmt.Println(nome)
	// verdade := false
	// fmt.Println(verdade)

	//CONSTANTES
	// não pode atribuir outro valor
	const idade = 30
	fmt.Println(idade)
	// não precisa ficar colocando "int/ string", o Go se entende
}
