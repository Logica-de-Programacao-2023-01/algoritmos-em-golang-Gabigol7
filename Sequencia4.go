package main

import "fmt"

func main() {
	var num1, num2, num3 float64

	fmt.Println("Digite o primeiro número:")
	fmt.Scanln(&num1)

	fmt.Println("Digite o segundo número:")
	fmt.Scanln(&num2)

	fmt.Println("Digite o terceiro número:")
	fmt.Scanln(&num3)

	mediaPonderada := (num1*2 + num2*3 + num3*5) / (2 + 3 + 5)

	fmt.Printf("A média ponderada é: %.2f\n", mediaPonderada)
}
