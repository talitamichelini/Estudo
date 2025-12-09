package main

// função = bloco de código que pode ser reutilizado
import "fmt"

func main() {
	// soma de 42 + 13
	// passando o mouse sobre"soma", ela mostra como abrir a função
	//fmt.Println(soma(42, 13))
	nome, sobrenome := printaNome("Talita", "Michelini")
	fmt.Println(nome)
	fmt.Println(sobrenome)
}

// func main é a função principal,
// o que de fato vai acontecer com nosso programa

// então é criada a função soma (exemplo) com os parâmetros declarados
// Pode ser soma, subtração, multiplicação, etc
//func soma(x int, y int) int {
// outro método de uso
// SomaDosValores := x + y
// return SomaDosValores
//	return x + y
//}

// sempre os parametros devem ser acompanhados do tipo
// e a função deve retornar um valor do tipo declarado
// nesse caso a função soma retorna um valor int
// quero resultado-retorno int, então declaro int após os parênteses

// para chamar a função soma, basta chamá-la dentro do main
// e passar os valores que queremos somar
// no caso 42 e 13

// usando exemplo de função associando a variável
// sempre iniciando com a função principal main

func printaNome(nome, sobrenome string) (string, string) {
	return nome, sobrenome
}

// função começando com letra minúscula é privada
// função começando com letra maiúscula é pública
