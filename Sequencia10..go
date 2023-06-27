package main

import "fmt"

func main() {
	var pesoKg float64

	// Leitura do peso em quilos
	fmt.Print("Digite o peso em quilos: ")
	fmt.Scan(&pesoKg)

	// Conversão para libras
	pesoLb := pesoKg * 2.20462

	// Exibição do resultado
	fmt.Printf("O peso em libras é: %.2f lb\n", pesoLb)
}
