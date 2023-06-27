package main

import "fmt"

func main() {
	var numero int

	fmt.Print("Digite um número inteiro: ")
	fmt.Scanln(&numero)

	antecessor := numero - 1
	sucessor := numero + 1

	fmt.Println("O antecessor é:", antecessor)
	fmt.Println("O sucessor é:", sucessor)
}
