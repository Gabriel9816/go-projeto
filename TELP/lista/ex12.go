package main

import "fmt"

func main() {
	// DETRATOR, Se a nota for 6 ,7  , o cliente é NEUTRO, se a nota for 9 ou 10 o cliente é PROMOTOR NPS- significa Net Promoter Score
	fmt.Println("Crie uma função  que classifique o  nível de satisfação do cliente usando METRICA NPS")
	var nota, nps float32
	var i, cliente int
	var DETRATORES, PASSIVOS, PROMOTORES float32 = 0, 0, 0
	var r1, r2 float32

	fmt.Printf("Digite quantos clientes vc quer receber avaliação\n\n")
	fmt.Scanf("%f", &cliente)

	//Entrada de dados
	for i = 1; i <= cliente; i++ {
		fmt.Printf("Digite a sua nota %d\n", i)
		fmt.Scanf("%f", &nota)

		if nota >= 0 && nota <= 6 {
			DETRATORES = DETRATORES + 1
		}
		if nota >= 7 && nota <= 8 {
			PASSIVOS = PASSIVOS + 1
		}
		if nota >= 9 && nota <= 10 {
			PROMOTORES = PROMOTORES + 1
		}
	}
	r1 = PROMOTORES / 150 * 100
	r2 = DETRATORES / 150 * 100
	nps = r1 - r2
	fmt.Printf("Nota NPS: %f", nps)

}
