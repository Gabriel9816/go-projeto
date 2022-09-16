package main

import "fmt"

func main() {
	i := 1

	var pi *int = &i

	i++
	fmt.Printf("pointer = %d, %d \n", i, &i)

	*pi++
	fmt.Printf("pi = %d, %d \n", pi, &pi)

}
