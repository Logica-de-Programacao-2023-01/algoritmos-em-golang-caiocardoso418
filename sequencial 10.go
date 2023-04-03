package main

import "fmt"

func main() {
	var peso float64
	fmt.Print("informe seu peso em kilos")
	fmt.Scan(&peso)

	libras := peso * 2.205

	fmt.Println("seu peso em libras é", libras)
}
