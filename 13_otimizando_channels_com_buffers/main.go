package main

import (
	"fmt"
	"time"
)

func numeros(n chan<- int) { // a setinha simboliza que só podemos escrever no channel
	for i := 0; i < 10; i++ {
		n <- i
		fmt.Printf("Escrito no channel: %d\n", i)
		time.Sleep(time.Millisecond * 90)
	}
	fmt.Println("Fim da escrita")
	close(n) // fechar o channel de comunicação
}

func main() {
	//cn := make(chan int)
	cn := make(chan int, 3) // colocando um número após o chan, dizemos que esse channel tem um espaço de 3 valores, ou seja, podemos preencher com 3 valores dentro do mesmo channel
	go numeros(cn)

	for v := range cn {
		fmt.Printf("lido do channel: %d\n", v)
		time.Sleep(time.Millisecond * 180)
	}

	<-cn

	fmt.Println("Fim da execução!")
}
