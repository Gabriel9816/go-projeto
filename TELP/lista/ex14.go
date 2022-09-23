package main

import "fmt"

func main() {
	fmt.Println("crie uma função que calcula o IMC")

	var peso float32
	var altura float32

	fmt.Print("Digite seu peso: ")
	fmt.Scanln(&peso)
	fmt.Print("Digite sua altura: ")
	fmt.Scanln(&altura)
	fmt.Println(" ")
	var r2 float32 = altura * altura
	var r float32 = peso / r2
	fmt.Printf("seu imc: %.2f \n", r)
}
