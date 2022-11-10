package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

const floatPrecision = 1000000

func GetRandInt(min, max int) int {
	nBig, _ := rand.Int(rand.Reader, big.NewInt(int64(max+1-min)))
	n := nBig.Int64()
	return int(n) + min
}

func GetRandFloat(min, max float64) float64 {
	minInt := int(min * floatPrecision)
	maxInt := int(max * floatPrecision)

	return float64(GetRandInt(minInt, maxInt)) / floatPrecision
}
func main() {

	fmt.Printf("Questão 01/10 - Escreva um algoritmo em go usando array que calcule a média simples de 99 valores (.float64) \n")

	var array [99]float64
	for i := 0; i < 99; i++ {
		array[i] = GetRandFloat(1, 10)
	}

	n := 99
	var soma float64 = 0
	for i := 0; i < n; i++ {
		soma += (array[i])
	}
	media := (float64(soma)) / (float64(n))
	fmt.Printf("\nMedia = %.2f \n", media)
}
