package main

import "fmt"

func main() {
	var num1, num2 int

	fmt.Print("Digite o primeiro número inteiro: ")
	fmt.Scan(&num1)

	fmt.Print("Digite o segundo número inteiro: ")
	fmt.Scan(&num2)

	if num1 > num2 {
		fmt.Println("O maior número é:", num1)
	} else if num2 > num1 {
		fmt.Println("O maior número é:", num2)
	} else {
		fmt.Println("Os números são iguais.")
	}
}
