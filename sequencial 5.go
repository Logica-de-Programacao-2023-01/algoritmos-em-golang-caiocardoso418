package main

import "fmt"

func main() {
	var idade float64
	fmt.Print("qual sua idade")
	fmt.Scan(&idade)

	dias := idade * 365

	fmt.Println("sua idade em dias é", dias)

}
