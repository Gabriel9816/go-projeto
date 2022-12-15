package main

import "fmt"

func main() {

	fmt.Printf("Questão 02/10 - Escreva um algoritmo em go usando array contendo 5 valores inteiros de 8bits que calcule um array de saida .float64" +
		"com valores normalizados 0 a 1 de saida e escreva o resultado")

	var inputArray [5]int8

	inputArray[0] = 10
	inputArray[1] = 20
	inputArray[2] = 30
	inputArray[3] = 40
	inputArray[4] = 50

	var outputArray [5]float64

	for i := 0; i < len(inputArray); i++ {
		outputArray[i] = float64(inputArray[i]) / 255.0
	}

	fmt.Println("\n", outputArray)

}
