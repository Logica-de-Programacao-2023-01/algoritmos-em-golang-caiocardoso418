package main

import "fmt"

func main() {
	var dias int
	var diaria float64

	fmt.Print("informe seus dias")
	fmt.Scan(&dias)

	fmt.Print("informe sua diaria")
	fmt.Scan(&diaria)

	salario := dias * dias

	fmt.Println("seu salario vai ser de", salario)
}
