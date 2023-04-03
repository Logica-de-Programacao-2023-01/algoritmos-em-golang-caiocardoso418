package main

import "fmt"

func main() {
	var num1, num2, num3 float64

	fmt.Print("informe os numeros")
	fmt.Scan(&num1, &num2, &num3)

	media := (2*num1 + 3*num2 + 5*num3) / 10

	fmt.Println("sua media ponderada é igual a", media)
}
