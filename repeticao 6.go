package main

import "fmt"

func main() {
	var num int
	fmt.Print("escreva um numero de 1 a 10: ")
	fmt.Scan(&num)

	for i := 1; i <= 10; i++ {
		resultado := num * i
		fmt.Println("a multiplicaçao é", resultado)
	}
}
