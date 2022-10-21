package main

import (
	"fmt"
	"sort"
)

func NamesByAlphabethicOrder(names []string) {
	fmt.Println("Escrever uma lista de nomes que retorne em ordem alfabética")
	fmt.Println("Lista de Nomes Original:")
	fmt.Println(names)
	sort.Strings(names)
	fmt.Println("Lista de Nomes Ordenada:")
	fmt.Println(names)
}
