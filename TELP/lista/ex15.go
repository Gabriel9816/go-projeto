package main

import "fmt"

func main() {
	fmt.Println("Faça um programa que mostra na tela os número de 1 a 100 e print em uma matriz quais deles sao numero de Ouro ?")
	fmt.Println("Números de 1 a 100")
	for n := 1; n <= 100; n++ {
		if n < 10 {
			fmt.Printf("0%d ", n)
		} else {
			fmt.Printf("%d ", n)
		}
		if n%20 == 0 {
			fmt.Printf("\n")
		}
	}
	fmt.Println("Números de Ouro")
	fmt.Println("[]")
}
