package main

import (
	"fmt"
)

// Escreva um programa que receba duas
//strings e verifique se elas são anagramas.
//Informe ao usuário se são ou não.

func main() {
	var x, y string
	fmt.Print("Escreva algo: ")
	fmt.Scan(&x)
	fmt.Print("Escreva algo: ")
	fmt.Scan(&y)

	if len(x) != len(y) {
		fmt.Println("Não é um anagrama")
	} else {
		frequenciaX := make(map[rune]int)
		frequenciaY := make(map[rune]int)

		for _, letra := range x {
			frequenciaX[letra] += 1
		}

		for _, letra := range y {
			frequenciaY[letra] += 1
		}

		anagrama := true

		for letra, count := range frequenciaX {
			if frequenciaY[letra] != count {
				anagrama = false
			}
		}

		if anagrama {
			fmt.Printf("As palavras %s e %s são anagramas\n", x, y)
		} else {
			fmt.Println("Não são anagramas")
		}
	}
}
