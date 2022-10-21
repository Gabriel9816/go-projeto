package main

import (
	"fmt"
	"time"
)

func BirthWeekDay(birth_day int, birth_month int, birth_year int) {
	fmt.Println("Escrever algoritmo em Go que receba a data de nascimento do usuário como entrada,  e retorne o dia da semana que o usuário nasceu.")
	if birth_day < 1 || birth_day > 31 {
		fmt.Println("Valor inválido! o dia de nascimento deve estar entre 1 e 31")
		return
	}
	if birth_month < 1 || birth_month > 12 {
		fmt.Println("Valor inválido! o mês de nascimento deve estar entre 1 e 12")
		return
	}
	if birth_year < 1 || birth_year > 2023 {
		fmt.Println("Valor inválido! o ano de nascimento deve estar entre 1 e 2023")
		return
	}

	var month = time.Month(birth_month)
	var date = time.Date(birth_year, month, birth_day, 23, 0, 0, 0, time.UTC)
	var week_day_number = int(date.Weekday())
	var week_day string

	switch week_day_number {
		case 0: week_day = "Segunda-feira"
		case 1: week_day = "Terça-feira"
		case 2: week_day = "Quarta-feira"
		case 3: week_day = "Quinta-feira"
		case 4: week_day = "Sexta-feira"
		case 5: week_day = "Sábado"
		case 6: week_day = "Domingo"
	}

	switch {
		case birth_day < 10 && birth_month < 10: fmt.Printf("A dia da data 0%d/0%d/%d corresponde a %s\n", birth_day, birth_month, birth_year, week_day)
		case birth_day < 10 && birth_month >= 10: fmt.Printf("A dia da data 0%d/%d/%d corresponde a %s\n", birth_day, birth_month, birth_year, week_day)
		case birth_day >= 10 && birth_month < 10: fmt.Printf("A dia da data %d/0%d/%d corresponde a %s\n", birth_day, birth_month, birth_year, week_day)
		case birth_day >= 10 && birth_month >= 10: fmt.Printf("A dia da data %d/%d/%d corresponde a %s\n", birth_day, birth_month, birth_year, week_day)
	}
}
