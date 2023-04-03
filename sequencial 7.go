package main

import "fmt"

func main() {
	var num float64
	fmt.Print("informe seu numero")
	fmt.Scan(&num)

	sucessor := num + 1
	antecessor := num - 1

	fmt.Println("seu sucessor é", sucessor)
	fmt.Println("seu antecessor é", antecessor)

}
