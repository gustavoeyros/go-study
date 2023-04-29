package main

import (
	"fmt"
	"time"
)

func numeros(done chan<- bool) { // a setinha simboliza que só podemos escrever no channel
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
		time.Sleep(time.Millisecond * 150)
	}
	done <- true // channels tem uma maneira diferente de receber valores
}

func letras(done chan bool) {
	for l := 'a'; l < 'j'; l++ {
		fmt.Printf("%c ", l)
		time.Sleep(time.Millisecond * 230)
	}
	done <- true
}

func main() {
	cn := make(chan bool)
	cl := make(chan bool)

	go numeros(cn) // para habilitar as goroutines, basta colocar a palavra "go" antes da função. Com isso, a função foi executada em modo de concorrência.
	go letras(cl)

	// a goroutine principal fica lendo o channel cn e cl e quando não receber mais valor, ela executa o fmt
	<-cn
	<-cl

	fmt.Println("Fim da execução!")
}
