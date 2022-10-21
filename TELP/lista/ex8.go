package main

import "fmt"

func RemoveVowels(phrase string) {
	fmt.Println("Crie uma função que retira as vogais de uma frase")

	var utf8_phrase []rune = []rune(phrase)
	var new_phrase string

	for i := 0; i < len(utf8_phrase); i++ {
		switch utf8_phrase[i] {
		case 65, 69, 73, 79, 85:
		case 97, 101, 105, 111, 117:
		case 192, 193, 194, 195, 201, 202, 205, 211, 212, 213, 218:
		case 224, 225, 226, 227, 233, 234, 237, 243, 244, 245, 250:
		default:
			// fmt.Printf("%c-(%d)___", utf8_phrase[i], utf8_phrase[i])
			new_phrase += string(utf8_phrase[i])
		}
	}

	fmt.Printf("Frase Original: %s.  Frase sem Vogais: %s\n", phrase, new_phrase)
}
