package main

import "fmt"

func main() {
	var (
		numero, soma, contador int
	)

	fmt.Println("Digite os números inteiros (digite 0 para sair):")

	for {
		fmt.Print(">> ")
		fmt.Scan(&numero)

		if numero == 0 {
			break
		}

		soma += numero
		contador++
	}

	if contador == 0 {
		fmt.Println("Nenhum número foi inserido.")
	} else {
		media := float64(soma) / float64(contador)
		fmt.Printf("A média aritmética é %.2f\n", media)
	}
}
