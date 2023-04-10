package main

import "fmt"

func main() {
	var idade int

	fmt.Println("Digite a idade da pessoa: ")
	fmt.Scan(&idade)

	if idade <= 9 {
		fmt.Println("Mirim")
	} else if idade <= 13 {
		fmt.Println("Infantil")
	} else if idade <= 17 {
		fmt.Println("Juvenil")
	} else {
		fmt.Println("Adulto")
	}
}
