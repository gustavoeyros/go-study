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

func (p Pessoa) String() string {
	return fmt.Sprintf("Olá, meu nome é %s e eu tenho %d anos.", p.Nome, p.Idade)
}

func (c Categoria) HasParent() bool {
	return c.Pai != nil
}

func (c *Categoria) SetPai(pai *Categoria) {
	c.Pai = pai
}
func main() {

	p := Pessoa{"Gustavo", "Eyros", 18, true, "000"}
	fmt.Println(p)

	cat := Categoria{Nome: "Categoria 1"}
	cat.SetPai((&Categoria{Nome: "Pai"}))
	if !cat.HasParent() {
		fmt.Println("n tem pai")
	}

}
