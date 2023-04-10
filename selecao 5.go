package main

import "fmt"

func main() {
	var num int

	fmt.Print("escreva um numero: ")
	fmt.Scan(&num)

	if num%3 == 0 && num%5 == 0 {
		fmt.Println("seu numero é divisivel por 3 e 5")
	} else {
		fmt.Println("seu numermo nao é divisivel por 3 nem 5")
	}
}
