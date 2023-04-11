package main

import "fmt"

func main() {
	var num int
	fmt.Print("escreva um numero inteiro:")
	fmt.Scan(&num)

	for i := 1; i <= num; i++ {
		if num%i == 0 {
			fmt.Println("o resultado é", i)

		}

	}
}
