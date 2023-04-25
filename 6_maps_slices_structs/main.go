package main

import "fmt"

type Pessoa struct {
	Nome      string
	Sobrenome string
	Idade     uint8
	Status    bool
	cpf       string // com letra minúscula, ele se torna privado e não pode ser acessado de um package diferente
}

type Categoria struct {
	Nome string
	Pai  *Categoria // para utilizar uma struct dentro da própria struct, precisa utilizar o *
}

func main() {
	// ######### SLICES ################
	/* 	nomes := []string{"nome1", "nome2", "nome3"}
	   	nomes = append(nomes, "nome4")             //adicionar dado na lista
	   	fmt.Println(nomes, len(nomes), cap(nomes)) // o go por baixo dos panos não aloca somente mais 1 espaço, ele separa um valor um pouco maior na memória para otimizar futuros append's.
	   	fmt.Println(nomes)

	   	animais := make([]string, 10) // inicializar um slice vazio para popular posteriormente
	*/

	// ######### MAPS ################
	/* idades := make(map[string]uint8) // maps armazenam um conjunto de pares chave-valor e não possui ordenação.
	idades["Gustavo"] = 18
	idades["Nome2"] = 25
	idades["Nome3"] = 20

	fmt.Println(idades) */

	p := Pessoa{
		Nome:      "Gustavo",
		Sobrenome: "Eyros",
		Idade:     18,
		Status:    true,
	}

	fmt.Println(p.Nome, p.Idade)

}
