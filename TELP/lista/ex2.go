package main

import "fmt"

func main() {
	fmt.Println("Criar uma função que converta segundos para minutos e segundos. Ex: 80seg = 1min e 20seg")

	var segundos, m, s, resto int

	fmt.Printf("Digite uma quantidade de segundos: ")
	fmt.Scanf("%d", &segundos)

	resto = segundos % 3600
	m = resto / 60
	s = resto % 60
	fmt.Printf("%d seg = %d min e %d seg\n", segundos, m, s)
}
