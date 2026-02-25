package main

import (
	"fmt"
)

// Loops - Laços de repetição
// Repetir tarefas x vezes
// FOR é o único laço de repetição em Go
// Só tem o FOR
// SEMPRE que houver lista, percorrer com loop = FOR RANGE

func main() {

	// sum := 0

	// // i++ soma +1 no final, i = 1+1
	// // i-- subtrai -1 no final
	// for i := 0; i < 10; i++ {
	// 	// for inicialização; condição; pós {
	// 	// começa em 0
	// 	// roda e para até número menor que 10, ou seja, 9
	// 	// e soma +1 no final de cada iteração
	// 	// exemplo:
	// 	// sum 0 + i 0 = 0 SOMA 1, = 1
	// 	// sum 0 + i 1 = 1 SOMA 1, = 2
	// 	// sum 1 + i 2 = 3 SOMA 1, = 4
	// 	// ...
	// 	// o resultado disso o programa vai trazer
	// 	fmt.Println("Sum:", sum)
	// 	fmt.Println("Índice:", i)
	// 	// print se quiser visualizar o processo de soma
	// 	sum += i
	// 	// sum += é sum + i
	// 	// sum -= é sum - i
	// 	// sum *= é sum * i
	// 	// sum /= é sum / i
	// 	// sum %= é sum % i
	// }
	// // o que for criado dentro do for, só existe dentro do for

	// fmt.Println(sum)

	// LOOP INFINITO - PERIGOSO
	// Interromper com CTRL + C
	// for {
	// 	fmt.Println("MEU")
	// 	time.Sleep(2 * time.Second)
	// 	fmt.Println("ORGULHO")
	// 	time.Sleep(2 * time.Second)
	// 	// caso queira adicionar pausa de 2 ou mais segundos
	// }

	// FOR RANGE - percorrer listas, arrays, slices, maps, strings
	frutas := []string{"laranja", "maca", "banana", "uva"}
	for _, lista := range frutas {
		fmt.Println(lista)
		// _ esconde o índice
		// i mostra o índice
	}

}
