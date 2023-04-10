package main

import (
	"fmt"
	"strings"
)

// Escreva um programa que receba uma
// string e um padrão (outro string)
// e retorne todos os índices em que o
// padrão ocorre na string.
// Informe ao usuário o resultado.

func main() {
	var value, padrao string
	fmt.Print("Escreva algo: ")
	fmt.Scan(&value)
	fmt.Print("Escreva algo: ")
	fmt.Scan(&padrao)

	var indexes []int

	i := strings.LastIndex(value, padrao)
	for i != -1 {
		indexes = append([]int{i}, indexes...)
		value = value[:i]
		i = strings.LastIndex(value, padrao)
	}

	fmt.Println(indexes)
}
