package main

import "fmt"

func main() {
	var num1, num2, num3 int

	// Leitura dos números
	fmt.Print("Digite o primeiro número: ")
	fmt.Scanln(&num1)

	fmt.Print("Digite o segundo número: ")
	fmt.Scanln(&num2)

	fmt.Print("Digite o terceiro número: ")
	fmt.Scanln(&num3)

	// Verificação do menor número
	menor := num1

	if num2 < menor {
		menor = num2
	}

	if num3 < menor {
		menor = num3
	}

	// Exibição do menor número
	fmt.Printf("O menor número é: %d\n", menor)
}
