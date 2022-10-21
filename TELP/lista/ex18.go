package main

import "fmt"

func CaesarCipher(key int, text string)  {
	fmt.Println("A Cifra de César é a técnica mais simples usada pela criptografia. Usada por Júlio César para \n",
		"a transmissão de mensagens secretas,a cifra é baseada na troca de uma letra por outra sempre \n",
		"obedecendo um código (ou melhor, uma chave). Faça um programa com base no que leu a cima que é criptografia utilizando da Cifra de César.")
		var new_text []rune = []rune(text)
		var alphabet []rune = []rune("abcdefghijklmnopqrstuvwxyzàáãâéêóôõíúçABCDEFGHIJKLMNOPQRSTUVWXYZÀÁÃÂÉÊÓÕÍÚÇ")
		var cripted_text string
		var index int
	
		for i := 0; i < len(new_text); i++ {
			for pos, char2 := range alphabet {
				if new_text[i] == char2 {
					index = pos + key
				}
				if (index > len(text)) {
					index = index % len(alphabet)
				}
			}
			cripted_text += string(alphabet[index])
		}	
		
		fmt.Println("Texto Original:")
		fmt.Println(text)
		fmt.Println("Texto Cifrado:")
		fmt.Println(cripted_text)
	}