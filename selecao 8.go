package main

import "fmt"

func main() {
	var num int

	fmt.Print("Digite um número entre 1 e 7: ")
	fmt.Scan(&num)

	if num == 1 {
		fmt.Println("Domingo")
	} else if num == 2 {
		fmt.Println("Segunda-feira")
	} else if num == 3 {
		fmt.Println("Terça-feira")
	} else if num == 4 {
		fmt.Println("Quarta-feira")
	} else if num == 5 {
		fmt.Println("Quinta-feira")
	} else if num == 6 {
		fmt.Println("Sexta-feira")
	} else if num == 7 {
		fmt.Println("Sábado")
	} else {
		fmt.Println("Número inválido")
	}
}
