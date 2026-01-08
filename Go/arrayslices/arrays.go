package main

import "fmt"

func main() {
	// var array [2]string    // 2 elementos tipo string (fixo)
	// array[0] = "Talita"    // índice 0
	// array[1] = "Michelini" // índice 1
	//fmt.Println(array[0])
	//fmt.Println(array[1])
	//fmt.Println(array)

	//numPrimos := [5]int{2, 3, 5, 7, 11} // tamanho 5, tipo int
	// fmt.Println(numPrimos)
	// fmt.Println(numPrimos[0:4]) // fatiamento do índice 0 ao 3, porque pega até antes do 4,não inclui
	// fmt.Println(numPrimos[1:])  // do índice 1 até o final
	// fmt.Println(numPrimos[:3])  // do início até o índice 2
	// fmt.Println(numPrimos[:])   // do início ao fim
	// fmt.Println(numPrimos[1])   // índice 1

	slice := make([]string, 5)
	slice[0] = "Talita"
	slice[1] = "Michelini"
	fmt.Println(slice)
	fmt.Println(len(slice))
	// append = adicionar elementos
	slice = append(slice, "Novo elemento")
}

// Listas
// arrays/ slices = homogêneos
// todos os elementos tem o mesmo tipo [int ou string ou float]

// ARRAYS = tamanho fixo = não cresce nem diminui
// precisa definir o tamanho antes = quantidade de elementos
// número dentro dos colchetes []
// exemplo: var array [5]int
// exemplo: var array [3]string
// acessar elementos: array[0], array[1], array[2]

// exemplo: var array [5]int = [5]int{1, 2, 3, 4, 5}
// tamanho 5, tipo int, elementos 1,2,3,4,5. índices 0,1,2,3,4

// exemplo: var array = [3]string{"A", "B", "C"}

// SLICES = tamanho flexível (mais usado)
// também são arrays, elementos do mesmo tipo,
// porém mais dinâmico
// significa que não precisa definir o tamanho antes
// pode começar zerado e ir adicionando elementos
// não tem número dentro dos colchetes []

// tamanho dado pela função len() {length}
// índice = tamanho - 1 - começa em 0

// Maps = heterogêneos
// [key]value
// chave/valor
// chave = string - valor = qualquer tipo
// exemplo: map[string]int
// exemplo: map[string]string
// exemplo: map[string]float64
// exemplo: map[string]bool
