package main

import "fmt"

func main() {
	var peso float64
	var altura float64

	// Leitura do peso e altura
	fmt.Print("Digite o peso em quilogramas: ")
	fmt.Scanln(&peso)

	fmt.Print("Digite a altura em metros: ")
	fmt.Scanln(&altura)

	// Cálculo do IMC
	imc := peso / (altura * altura)

	// Exibição do resultado
	fmt.Printf("O IMC é: %.2f\n", imc)
}
