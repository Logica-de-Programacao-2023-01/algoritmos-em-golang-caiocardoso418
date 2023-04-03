package main

import "fmt"

func main() {
	var produto float64
	fmt.Print("infome o valor do produto")
	fmt.Scan(&produto)

	desconto := produto * 10 / 100
	resultado := produto - desconto

	fmt.Println("o valor do desconto vai ficar", resultado)

}
