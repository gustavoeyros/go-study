package main

import "fmt"

var nome string // Declaração de variável à nível de package, ou seja, podemos utilizar essa variável em outro arquivo do mesmo package.

var (
	numero1 int
	numero2 int
) // Conseguimos declarar diversas variáveis à nível de package.

func main() {
	mensagem := "Alguma mensagem" // Declara e atribui um valor ao mesmo tempo.
	fmt.Println(mensagem)

	var a, b, c = true, 2.5, "Eae!" // declarando e atribuindo vários valores

	fmt.Println(a, b, c)

	nome = "Gustavo Eyros"
	fmt.Println("Opa,", nome)

	// mudando valores de variáveis
	var x, y = 10, 20
	fmt.Println(x, y)

	x, y = y, x
	fmt.Println(x, y)
}
