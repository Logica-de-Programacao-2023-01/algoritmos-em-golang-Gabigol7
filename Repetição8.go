package main

import "fmt"

func main() {
	var numero int
	fmt.Print("Digite um número inteiro positivo: ")
	fmt.Scan(&numero)

	fmt.Printf("Divisores de %d: ", numero)
	for i := 1; i <= numero; i++ {
		if numero%i == 0 {
			fmt.Printf("%d ", i)
		}
	}
}
