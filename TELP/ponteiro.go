package main

import "fmt"

func main() {
	i := 1

	var pi *int = &i

	i++
	fmt.Printf("pointer = %d", i)

	*pi++
	fmt.Printf("pi = %d", pi)

}
