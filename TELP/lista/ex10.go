package main

import "fmt"

func main() {
	fmt.Println("uma função para dizer se o numero inserido é primo ou n")
	var num, i int
	var resultado int = 0

	fmt.Printf("Digite um número: ")
	fmt.Scanf("%d", &num)

	for i = 2; i <= num/2; i++ {
		if num%i == 0 {
			resultado++
			break
		}
	}

	if resultado == 0 {
		fmt.Printf("%d é um número primo\n", num)
	} else {
		fmt.Printf("%d não é um número primo\n", num)
	}
}
