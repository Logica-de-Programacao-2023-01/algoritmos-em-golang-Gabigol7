package main

import "fmt"

func main() {
	var salario float64

	fmt.Print("Digite o salário do funcionário: ")
	fmt.Scanln(&salario)

	aumento := salario * 0.15
	novoSalario := salario + aumento

	fmt.Printf("O novo salário com aumento é: %.2f\n", novoSalario)
}
