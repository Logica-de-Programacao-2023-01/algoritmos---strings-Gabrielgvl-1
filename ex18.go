package main

import (
	"fmt"
	"strconv"
)

// Solicite ao usuário uma string e informe se
// ela contém somente números (0 a 9).

func main() {
	var value string
	fmt.Print("Escreva algo: ")
	fmt.Scan(&value)

	_, err := strconv.Atoi(value)
	if err != nil {
		fmt.Println("A string não contém apenas números")
	} else {
		fmt.Println("A string contém apenas números")
	}
}
