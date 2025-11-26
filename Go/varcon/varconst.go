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
	// const idade = 30
	// fmt.Println(idade)
	// idade = 31 -> dá erro, porque não da pra adicionar outro valor a uma constante
	// não precisa ficar colocando "int/ string", o Go se entende

	// ZERO VALUES
	// dessa maneira os valores saem zerados
	var x int
	var y string
	var z bool
	var w float64
	fmt.Printf("Inteiro: %d\n", x)
	// placeholder %d = decimal (inteiro)
	fmt.Printf("String: %s\n", y)
	// placeholder %s = string (texto)
	fmt.Printf("Bool: %t\n", z)
	// placeholder %t = booleano (true/false)
	fmt.Printf("Float: %f\n", w)
	// placeholder %f = float (decimal)
	// OU %v para qualquer tipo de valor
}
