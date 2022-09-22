package main

import "fmt"

func main() {
	fmt.Println("Criar uma função que faça média entre notas")
	var n1, n2 float32

	fmt.Print("Digite o n1: ")
	fmt.Scanln(&n1)
	fmt.Print("Digite o n2: ")
	fmt.Scanln(&n2)
	fmt.Println(" ")
	var r1 float32 = n1 + n2
	var r2 float32 = r1 / 2
	fmt.Printf("média: %.2f \n", r2)
}
