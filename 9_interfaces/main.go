package main

import "fmt"

type Pessoa struct {
	Nome   string
	Idade  uint8
	Status bool
}

type Document interface {
	Doc() string
}

func (p Pessoa) String() string {
	return fmt.Sprintf("Olá, meu nome é %s e eu tenho %d anos.", p.Nome, p.Idade)
}

func (pf PessoaFisica) Doc() string {
	return fmt.Sprintf("Meu CPF é  %s", pf.cpf)
}

type PessoaFisica struct {
	Pessoa    // em go, ao invés de usarmos extends para herança, basta colocar o nome da outra struct
	Sobrenome string
	cpf       string
}

type PessoaJuridica struct {
	Pessoa
	RazaoSocial string
	cnpj        string
}

func (pj PessoaJuridica) Doc() string {
	return fmt.Sprintf("meu cnpj é %s", pj.cnpj)
}

func show(d Document) {
	fmt.Println(d)
	fmt.Println(d.Doc())
}

func main() {
	pf := PessoaFisica{
		Pessoa{Nome: "Gustavo", Idade: 18, Status: true},
		"Eyros",
		"000.000.000-00",
	}

	show(pf)

	pj := PessoaJuridica{
		Pessoa{Nome: "Golangzada", Idade: 1, Status: true},
		"Sla LTDA",
		"000.000.000-00",
	}

	show(pj)
}
