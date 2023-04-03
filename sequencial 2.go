package main

import "fmt"

func main() {
	var num1 uint
	fmt.Print("escolha um numero")
	fmt.Scan(&num1)

	dobro := num1 * 2
	triplo := num1 * 3
	quadroplo := num1 * 4

	fmt.Println("o dobro é", dobro)
	fmt.Println("o triplo é", triplo)
	fmt.Println("o quadroplo é", quadroplo)

}
