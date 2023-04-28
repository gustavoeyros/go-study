package main

import "fmt"

type Pessoa struct {
	Nome   string
	Idade  uint8
	Status bool
}

func (p Pessoa) String() string {
	return fmt.Sprintf("Olá, meu nome é %s e eu tenho %d anos.", p.Nome, p.Idade)
}

type PessoaFisica struct {
	Pessoa    // em go, ao invés de usarmos extends para herança, basta colocar o nome da outra struct
	Sobrenome string
	cpf       string
}

type PessoaJuridica struct {
	RazaoSocial string
	cnpj        string
}

func main() {
	p := PessoaFisica{
		Pessoa{Nome: "Gustavo", Idade: 18, Status: true},
		"Eyros",
		"000.000.000-00",
	}

	fmt.Println(p)
}
