package main

import "fmt"

func main() {
	fmt.Println("Números ímpares de 1 a 19:")
	for i := 1; i <= 19; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}
}
