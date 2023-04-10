package main

import "fmt"

func main() {
	var num1 float64
	var num2 float64
	var num3 float64

	fmt.Print("informe o primeiro numeor")
	fmt.Scan(&num1)
	fmt.Print("informe o segundo numero")
	fmt.Scan(&num2)
	fmt.Print("infome o terceiro numro")
	fmt.Scan(&num3)

	if num1 < num2 && num1 < num3 {
		fmt.Println("o primeiro numero é o menor")
	} else if num2 < num3 && num2 < num1 {
		fmt.Println("o segundo numero é o menor")
	} else if num3 < num2 && num3 < num1 {
		fmt.Println("o terceiro numero é o menor")
	} else {
		fmt.Println("os numeros sao iguais")
	}
}
