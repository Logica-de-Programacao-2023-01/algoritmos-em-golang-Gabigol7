package main

import "fmt"

func main() {
	var num, max int

	// Lê o primeiro número
	fmt.Print("Digite um número inteiro (ou 0 para sair): ")
	fmt.Scan(&num)
	max = num // define o primeiro número como o maior até agora

	// Loop para a leitura dos números
	for num != 0 {
		// Lê o próximo número
		fmt.Print("Digite um número inteiro (ou 0 para sair): ")
		fmt.Scan(&num)

		// Verifica se o número atual é maior que o máximo
		if num > max {
			max = num
		}
	}

	// Exibe o maior número
	fmt.Println("O maior número digitado foi:", max)
}
