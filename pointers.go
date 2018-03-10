package main

import "fmt"

func main() {
	something := 10
	addrOfSomething := &something
	// addrOfSomething++ // Erro! Não é um inteiro!
	*addrOfSomething++
	fmt.Printf("Valor: %v\nEndereço: %v\n", *addrOfSomething, addrOfSomething)
}