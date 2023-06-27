package main

import "fmt"

func main() {
	var idade int

	fmt.Print("Digite a idade da pessoa: ")
	fmt.Scan(&idade)

	var classificacao string

	if idade <= 9 {
		classificacao = "mirim"
	} else if idade <= 13 {
		classificacao = "infantil"
	} else if idade <= 17 {
		classificacao = "juvenil"
	} else {
		classificacao = "adulto"
	}

	fmt.Println("A classificação da pessoa é:", classificacao)
}
