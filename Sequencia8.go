package main

import "fmt"

func main() {
	var diasTrabalhados int
	var valorDiaria float64

	// Leitura do número de dias trabalhados e valor da diária
	fmt.Print("Digite o número de dias trabalhados: ")
	fmt.Scanln(&diasTrabalhados)

	fmt.Print("Digite o valor da diária: ")
	fmt.Scanln(&valorDiaria)

	// Cálculo do salário
	salario := float64(diasTrabalhados) * valorDiaria

	// Exibição do resultado
	fmt.Printf("O salário do funcionário é: R$ %.2f\n", salario)
}
