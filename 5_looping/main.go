package main

import "fmt"

func main() {

	nomes := []string{"Nome 1", "Nome 2", "Nome 3"}

	/* 	for i := 0; i < 10; i++ {
		fmt.Println(i)
	} */

	/* 	for i := 0; i < len(nomes); i++ {
		fmt.Println(nomes[i])
	} */

	/* for k, nome := range nomes {
		fmt.Println(k, nome)
	} */

	var i int
	for i < len(nomes) {
		fmt.Println(nomes[i])
		i++
	}

}
