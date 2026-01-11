package main

import "fmt"

// STRUCTS = forma de criar sua própria estrutura de dados
// personalizar de acordo com a necessidade do seu programa
// pode usar vários tipos diferentes

// type = palavra reservada para criar um novo tipo
// nome do tipo = nome que você quer dar para o seu novo tipo
// struct = palavra reservada para criar uma estrutura
// {} = chaves para definir o escopo da estrutura

// type - nome da estrutura - struct {
// 	 campos da estrutura (identado)
// }

type Pessoa struct {
	Nome string
	//Sobrenome string
	Idade int
	//AnoNasc int
	// possibilidade de criar vários campos
}

// criando outra struct
// exemplo herdando dados da struct Pessoa
type Profissao struct {
	Pessoa // nome + idade
	Tipo   string
	// segue o código no final
}

func main() {
	// fmt.Println(Pessoa{"Talita", 29})
	// fmt.Println(Pessoa{Nome: "Blair", Idade: 2})

	p1 := Pessoa{Nome: "Talita", Idade: 29}
	// // variável p1 do tipo Pessoa
	// fmt.Println(p1.Nome)
	// fmt.Println(p1.Idade)

	// p1.Idade = 30
	//fmt.Println(p1.Idade)
	// alterar valor do campo Idade

	p2 := Pessoa{Nome: "Gabrielle", Idade: 28}
	// variável p2 do tipo Pessoa
	// Pessoa é um Tipo que criei, leva os campos Nome e Idade
	// fmt.Println(p2.Nome)
	// fmt.Println(p2.Idade)

	p3 := Pessoa{Nome: "Blair", Idade: 2}
	// criando várias variáveis do tipo Pessoa

	// var pessoas [3]Pessoa = [3]Pessoa{p1, p2, p3}
	// tamanho 3
	pessoas := []Pessoa{}
	pessoas = append(pessoas, p1, p2, p3)
	fmt.Println(pessoas)
	// structs + slices

	// map com slice de struct
	// chave string - valor slice de Pessoa
	// alunos := map[string][]Pessoa{}
	// alunos["Equipe Home Office"] = pessoas
	// fmt.Println(alunos)
	// var - string - lista de pessoas
	// Outro método
	var alunos = map[string][]Pessoa{
		//	"Equipe Home Office": {{Nome: "Talita", Idade: 29}, {Nome: "Gabrielle", Idade: 28}},
		"Equipe Home Office": pessoas,
		// reutilizando a variável pessoas, da na mesma
	} // chave - valor
	fmt.Println(alunos)

	// usando a struct Profissao
	profissa := Profissao{p3, "Cachorro"}
	fmt.Println(profissa)
	fmt.Println(profissa.Nome)  // printa nome isolado
	fmt.Println(profissa.Idade) // printa idade isolada
	fmt.Println(profissa.Tipo)  // printa tipo isolado
}

// métodos e exemplos de uso
