package main

import "fmt"

func main() {
	var salario float64

	fmt.Println("Digite o salário do funcionário:")
	fmt.Scanln(&salario)

	var novoSalario float64

	if salario < 1000.00 {
		novoSalario = salario + (salario * 0.10)
	} else {
		novoSalario = salario + (salario * 0.05)
	}

	fmt.Println("Novo salário:", novoSalario)
}
