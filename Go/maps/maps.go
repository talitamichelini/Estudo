package main

import "fmt"

func main() {
	// idade := map[string]int{}
	// idade["Talita"] = 29
	// fmt.Println(idade)

	anoNascimento := map[string]int{
		"Talita": 1996}
	fmt.Println(anoNascimento)
	fmt.Println(anoNascimento["Talita"])
	anoNascimento["Alves"] = 29
	fmt.Println(anoNascimento)
}

// MAPS - heterogeneos
// pode misturar tipos
// [key]value
// chave um tipo, valor outro tipo

// map[string]int
// "apple":  5,
//  "banana": 10,
//  "orange": 7
// map serve para representar dicionarios ou objetos
