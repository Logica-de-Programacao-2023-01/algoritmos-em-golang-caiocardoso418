package main

import "fmt"

func main() {
	var num1 int
	var num2 int

	fmt.Println("escreva um numero:")
	fmt.Scan(&num1)
	fmt.Print("escreva um numero:")
	fmt.Scan(&num2)

	multiplicacao := num1 * num2
	soma := num1 + num2

	if num1 >= 0 && num2 >= 0 {
		fmt.Println("o resultado da multiplicação é:", multiplicacao)
	} else {
		fmt.Println("o resultado da soma é:", soma)
	}
}
