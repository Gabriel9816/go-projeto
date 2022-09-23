package main

import "fmt"

func main() {
	fmt.Println("Criar uma função que converta a moeda Real , para moeda dólar")
	var real float32
	var dolar float32 = 5.12

	fmt.Print("Digite o valor em real: ")
	fmt.Scanln(&real)
	fmt.Println(" ")
	var r float32 = real / dolar
	fmt.Printf("real em dolar: %.2f \n", r)
}
