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

func (c Categoria) HasParent() bool {
	return c.Pai != nil
}

func main() {

	cat := Categoria{Nome: "Categoria 1"}
	if !cat.HasParent() {
		fmt.Println("n tem pai")
	}

}
