package main

import "fmt"

func main() {
	var num1, num2, num3 float64

	fmt.Println("Digite três números:")
	fmt.Scan(&num1, &num2, &num3)

	if num1 > num2 {
		num1, num2 = num2, num1
	}
	if num1 > num3 {
		num1, num3 = num3, num1
	}
	if num2 > num3 {
		num2, num3 = num3, num2
	}

	fmt.Println("Os números em ordem crescente são:", num1, num2, num3)
}
