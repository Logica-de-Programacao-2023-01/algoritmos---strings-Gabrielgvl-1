package main

import (
	"fmt"
	"strings"
)

// Solicite ao usuário uma string e
//informe se ela é está em camelCase e
//quantas palavras possuí.
//CamelCase é quando a primeira letra de
//cada palavra é maiúscula,
//e as demais são minúsculas.
//Exemplo: "estaStringEstaEmCamelCase".

func main() {
	var value string
	fmt.Print("Escreva algo: ")
	fmt.Scan(&value)

	if strings.ToLower(string(value[0])) != string(value[0]) {
		fmt.Println("Não está em camelCase")
	} else {
		palavras := 1
		for _, x := range value {
			if strings.ToUpper(string(x)) == string(x) {
				palavras++
			}
		}
		fmt.Printf("Tem %d palavras na string\n", palavras)
	}
}
