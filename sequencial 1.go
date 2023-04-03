package main

import "fmt"

func main() {
	var num1 uint
	fmt.Print("qual o numero")
	fmt.Scan(&num1)

	var num2 uint
	fmt.Print("qual o numero")
	fmt.Scan(&num2)

	var num3 uint
	fmt.Println("qual o numero")
	fmt.Scan(&num3)

	fmt.Println("a soma entre os 3 numeros é", num1+num3+num2)

}
