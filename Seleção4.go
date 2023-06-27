package main

import "fmt"

func main() {
	var altura float64
	var sexo string

	fmt.Print("Digite a altura (em metros): ")
	fmt.Scanln(&altura)

	fmt.Print("Digite o sexo (M para masculino ou F para feminino): ")
	fmt.Scanln(&sexo)

	var pesoIdeal float64
	if sexo == "M" {
		pesoIdeal = (72.7 * altura) - 58
	} else if sexo == "F" {
		pesoIdeal = (62.1 * altura) - 44.7
	} else {
		fmt.Println("Sexo inválido. Digite M para masculino ou F para feminino.")
		return
	}

	fmt.Println("O peso ideal para a altura fornecida é:", pesoIdeal)

	var peso float64
	fmt.Print("Digite o peso da pessoa (em kg): ")
	fmt.Scanln(&peso)

	if peso < pesoIdeal {
		fmt.Println("A pessoa está abaixo do peso ideal.")
	} else if peso > pesoIdeal {
		fmt.Println("A pessoa está acima do peso ideal.")
	} else {
		fmt.Println("A pessoa está dentro do peso ideal.")
	}
}
