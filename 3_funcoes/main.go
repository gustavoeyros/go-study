package main

import (
	"fmt"
	"strconv"
)

// para deixar a função à nível de pacote, basta deixar a primeira letra da função em maiúsculo. Ex:
func Hello(nome string) {
	fmt.Println("Hello,", nome, "!")
}

func sum(a, b int) int {
	return a + b
}

func convertAndSum(a int, b string) (total int, err error) { // quando nomeamos as variáveis de retornos, elas já são inicializadas quando invocadas
	// i, _ := strconv.Atoi(b) // atoi converte o valor para inteiro, porém, essa função possui o parâmetro err, pra omitir ele, basta usar "_"
	i, err := strconv.Atoi(b)

	if err != nil {
		return
	}
	total = a + i
	return
}

func main() {
	Hello("Gustavo")

	fmt.Println("Total: ", sum(10, 5))
	total, err := convertAndSum(10, "25")
	fmt.Println("Total: ", total, err)
}
