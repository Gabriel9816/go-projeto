package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	var name string = "Samsung"
	var age int = 23
	//var n1, n2 float32 = 10.35, 20.53
	var n1, n2 float32
	var task bool = true
	fmt.Printf("Aluno: %s tem %d anos e passou no tads 6? %t\n", name, age, task)
	fmt.Println(" ")
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
