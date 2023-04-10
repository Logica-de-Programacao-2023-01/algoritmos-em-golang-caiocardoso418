package main

import "fmt"

func main() {
	var salario float64

	fmt.Println("Digite o salário:")
	fmt.Scan(&salario)

	if salario < 1000 {
		salario = salario * 1.1
	} else {
		salario = salario * 1.05
	}

	fmt.Println("O novo salário é", salario)
}
