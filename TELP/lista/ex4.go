package main

import "fmt"

func main() {
	fmt.Println("crie uma função que some duas notas")
	var n1, n2 float32
	
	fmt.Print("Digite o n1: ")
	fmt.Scanln(&n1)
	fmt.Print("Digite o n2: ")
	fmt.Scanln(&n2)
	fmt.Println(" ")
	var r1 float32 = n1 + n2

	fmt.Printf("Soma de n1 e n2: %.2f \n", r1)
}
