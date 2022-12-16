package main

import (
	"fmt"
)

func main() {

	fmt.Printf("Questão 03/10 - Mapa em go de 3 notas, imprimir os valores")
	
	notas := make(map[string]float64)

	notas["samsung"] = 7.5
	notas["caio"] = 8.5
	notas["mauricio"] = 9.0

	for nome, nota := range notas {
		fmt.Printf("\n %s tem nota %.1f\n", nome, nota)
	}

}
