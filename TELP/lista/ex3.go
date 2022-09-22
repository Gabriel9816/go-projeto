package main

import "fmt"

func main() {
	fmt.Println("Criar uma função que converta anos para minutos e segundos.")
	var Anos, Dias, Horas, Meses, Minutos, Segundo int

	fmt.Printf("Insira o numero de anos: ")
	fmt.Scanf("%d", &Anos)

	Meses = Anos * 12
	Dias = Anos * 365
	Horas = Dias * 24
	Minutos = Horas * 60
	Segundo = Minutos * 60

	fmt.Printf("\n%d anos equivalem a \n%d meses, \n%d dias, \n%d horas, \n%d minutos, \n%d segundos.\n", Anos, Meses, Dias, Horas, Minutos, Segundo)
}
