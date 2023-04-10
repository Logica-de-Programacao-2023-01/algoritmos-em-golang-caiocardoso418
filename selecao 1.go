package main

import "fmt"

func main() {

	var num1 float64
	var num2 float64

	fmt.Print("escreva um numero")
	fmt.Scan(&num1)

	fmt.Print("escreva um numero")
	fmt.Scan(&num2)

	if num1 > num2 {
		fmt.Println("num1 é o maior")
	} else {
		fmt.Println("num2 é o maior")
	}
}
