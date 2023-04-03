package main

import "fmt"

func main() {
	var salario float64

	fmt.Print("informe seu salário")
	fmt.Scan(&salario)

	aumento := (15 * salario) / 100
	resultado := aumento + salario

	fmt.Println("seu salario novo é", resultado)
}
