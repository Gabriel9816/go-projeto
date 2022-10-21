package main

import "fmt"

func main() {
	fmt.Println("Crie uma função recebe uma temperatura em graus celsius e converta para Fahrenheit.")
	var c float32

	fmt.Print("Digite a temperatura em Celcius: ")
	fmt.Scanln(&c)
	fmt.Println(" ")
	var f float32 = c*1.8 + 32 //formula para converter celsius em Fahrenheit
	fmt.Printf("Temperatura em Fahrenheit: %.2f \n", f)
}
