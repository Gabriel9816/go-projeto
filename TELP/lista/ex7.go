package main

import "fmt"

func main() {
	fmt.Println("Crie uma função que receba 10 números, faça a média e mostre o maior número")
	var n1, n2, n3, n4, n5, n6, n7, n8, n9, n10 float32

	fmt.Print("Digite o n1: ")
	fmt.Scanln(&n1)
	fmt.Print("Digite o n2: ")
	fmt.Scanln(&n2)
	fmt.Print("Digite o n3: ")
	fmt.Scanln(&n3)
	fmt.Print("Digite o n4: ")
	fmt.Scanln(&n4)
	fmt.Print("Digite o n5: ")
	fmt.Scanln(&n5)
	fmt.Print("Digite o n6: ")
	fmt.Scanln(&n6)
	fmt.Print("Digite o n7: ")
	fmt.Scanln(&n7)
	fmt.Print("Digite o n8: ")
	fmt.Scanln(&n8)
	fmt.Print("Digite o n9: ")
	fmt.Scanln(&n9)
	fmt.Print("Digite o n10: ")
	fmt.Scanln(&n10)
	fmt.Println(" ")
	var r1 float32 = n1 + n2 + n3 + n4 + n5 + n6 + n7 + n8 + n9 + n10
	var r2 float32 = r1 / 10
	fmt.Printf("média: %.2f \n", r2)
	if n1 > n2 || n1 > n3 || n1 > n4 || n1 > n5 || n1 > n6 || n1 > n7 || n1 > n8 || n1 > n9 || n1 > n10 {
		fmt.Printf("o maior numero é: %.2f \n", n1)
	}
	if n2 > n1 || n2 > n3 || n2 > n4 || n2 > n5 || n2 > n6 || n2 > n7 || n2 > n8 || n2 > n9 || n2 > n10 {
		fmt.Printf("o maior numero é: %.2f \n", n2)
	}
	if n3 > n2 || n3 > n1 || n3 > n4 || n3 > n5 || n3 > n6 || n3 > n7 || n3 > n8 || n3 > n9 || n3 > n10 {
		fmt.Printf("o maior numero é: %.2f \n", n3)
	}
	if n4 > n2 || n4 > n3 || n4 > n1 || n4 > n5 || n4 > n6 || n4 > n7 || n4 > n8 || n4 > n9 || n4 > n10 {
		fmt.Printf("o maior numero é: %.2f \n", n4)
	}
	if n5 > n2 || n5 > n3 || n5 > n4 || n5 > n1 || n5 > n6 || n5 > n7 || n5 > n8 || n5 > n9 || n5 > n10 {
		fmt.Printf("o maior numero é: %.2f \n", n5)
	}
	if n6 > n2 || n6 > n3 || n6 > n4 || n6 > n5 || n6 > n1 || n6 > n7 || n6 > n8 || n6 > n9 || n6 > n10 {
		fmt.Printf("o maior numero é: %.2f \n", n6)
	}
	if n7 > n2 || n7 > n3 || n7 > n4 || n7 > n5 || n7 > n6 || n7 > n1 || n7 > n8 || n7 > n9 || n7 > n10 {
		fmt.Printf("o maior numero é: %.2f \n", n7)
	}
	if n8 > n2 || n8 > n3 || n8 > n4 || n8 > n5 || n8 > n6 || n8 > n7 || n8 > n1 || n8 > n9 || n8 > n10 {
		fmt.Printf("o maior numero é: %.2f \n", n8)
	}
	if n9 > n2 || n9 > n3 || n9 > n4 || n9 > n5 || n9 > n6 || n9 > n7 || n9 > n8 || n9 > n1 || n9 > n10 {
		fmt.Printf("o maior numero é: %.2f \n", n9)
	}
	if n10 > n2 || n10 > n3 || n10 > n4 || n10 > n5 || n10 > n6 || n10 > n7 || n10 > n8 || n10 > n9 || n10 > n1 {
		fmt.Printf("o maior numero é: %.2f \n", n10)
	}
}
