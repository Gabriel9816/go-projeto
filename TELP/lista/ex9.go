package main

import "fmt"

func main() {
	fmt.Println("Receber três números e faça a multiplicação entre eles")
	var n1, n2, n3 float32

	fmt.Print("Digite o n1: ")
	fmt.Scanln(&n1)
	fmt.Print("Digite o n2: ")
	fmt.Scanln(&n2)
	fmt.Print("Digite o n3: ")
	fmt.Scanln(&n3)
	fmt.Println(" ")
	var r1 float32 = n1 * n2
	var r2 float32 = r1 * n3
	fmt.Printf("média: %.2f \n", r2)
}
