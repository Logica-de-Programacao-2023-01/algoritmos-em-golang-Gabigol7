package main

import "fmt"

func main() {
	var number int

	fmt.Print("Digite um número: ")
	fmt.Scan(&number)

	fmt.Println("Tabuada de multiplicação para", number)
	for i := 1; i <= 10; i++ {
		result := number * i
		fmt.Printf("%d x %d = %d\n", number, i, result)
	}
}
