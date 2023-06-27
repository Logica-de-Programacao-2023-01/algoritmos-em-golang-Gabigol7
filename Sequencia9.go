package main

import "fmt"

func main() {
	var preco float64

	fmt.Print("Digite o preço do produto: ")
	fmt.Scanln(&preco)

	desconto := preco * 0.1
	precoComDesconto := preco - desconto

	fmt.Printf("O valor com desconto de 10%% é: %.2f\n", precoComDesconto)
}
