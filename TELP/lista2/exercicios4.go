package main

import (
	"fmt"
	"unsafe"
)

func main() {

	fmt.Printf("Questão 04/10 - Slice com n posições, imprimir o endereço de memória da primeira posição e o conteúdo do slice em ordem decrescente;")

	n := 5
	slice := make([]int, n)

	fmt.Println("\nMemória: ", unsafe.Pointer(&slice[0]))

	for i := 0; i < len(slice); i++ {
		slice[i] = i
	}

	for i := len(slice) - 1; i >= 0; i-- {
		fmt.Println(slice[i])
	}
}
