package main

import "fmt"

func main() {
	func() {
		fmt.Println("Função anônima!")
	}()
	f := retornaFunção()
	f()
}

func retornaFunção() func() {
	return func() {
		fmt.Println("Função como valor.")
	}
}