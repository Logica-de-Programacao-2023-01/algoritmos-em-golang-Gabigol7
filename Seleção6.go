package main

import "fmt"

func main() {
	var num1, num2 int

	fmt.Println("Digite o primeiro número:")
	fmt.Scanln(&num1)

	fmt.Println("Digite o segundo número:")
	fmt.Scanln(&num2)

	if num1 >= 0 && num2 >= 0 {
		result := num1 * num2
		fmt.Println("Resultado da multiplicação:", result)
	} else {
		result := num1 + num2
		fmt.Println("Resultado da soma:", result)
	}
}
