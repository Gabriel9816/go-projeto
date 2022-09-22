package main

import "fmt"

func main() {
	fmt.Println("Uma função que cria uma calculadora simples, digitando o n1, n2 e a operação desejada")
	var n1, n2 float32
	fmt.Print("Digite o n1: ")
	fmt.Scanln(&n1)
	fmt.Print("Digite o n2: ")
	fmt.Scanln(&n2)
	fmt.Println(" ")
	var r1, r2, r3, r4 float32 = n1 + n2, n1 * n2, n1 / n2, n1 - n2
	var r5 float32 = r1 + r2 + r3 + r4

	fmt.Printf("Soma de n1 e n2: %.2f \n", r1)
	fmt.Printf("multiplicação de n1 e n2: %.2f \n", r2)
	fmt.Printf("Divisão de n1 e n2: %.2f \n", r3)
	fmt.Printf("Subtração de n1 e n2: %.2f \n", r4)
	fmt.Printf("Resultado de todas as operações: %.2f \n", r5)
}
