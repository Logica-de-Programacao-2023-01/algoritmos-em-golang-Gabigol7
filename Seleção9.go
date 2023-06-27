package main

import (
	"fmt"
	"sort"
)

func main() {
	var num1, num2, num3 float64

	fmt.Println("Digite o primeiro número:")
	fmt.Scanln(&num1)

	fmt.Println("Digite o segundo número:")
	fmt.Scanln(&num2)

	fmt.Println("Digite o terceiro número:")
	fmt.Scanln(&num3)

	// Coloca os números em um slice
	numeros := []float64{num1, num2, num3}

	// Ordena o slice em ordem crescente
	sort.Float64s(numeros)

	// Mostra os números em ordem crescente
	fmt.Println("Números em ordem crescente:", numeros)
}
