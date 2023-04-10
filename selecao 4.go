package main

import (
	"fmt"
)

func main() {
	var altura, peso float64
	var sexo string

	fmt.Print("Digite a altura (em metros): ")
	fmt.Scanln(&altura)
	fmt.Print("Digite o peso (em kg): ")
	fmt.Scanln(&peso)
	fmt.Print("Digite o sexo (M para masculino ou F para feminino): ")
	fmt.Scanln(&sexo)

	imc := peso / (altura * altura)

	var mensagem string

	if sexo == "M" {
		if imc < 20.7 {
			mensagem = "abaixo do peso ideal"
		} else if imc >= 20.7 && imc <= 26.4 {
			mensagem = "dentro do peso ideal"
		} else if imc > 26.4 {
			mensagem = "acima do peso ideal"
		}
	} else if sexo == "F" {
		if imc < 19.1 {
			mensagem = "abaixo do peso ideal"
		} else if imc >= 19.1 && imc <= 25.8 {
			mensagem = "dentro do peso ideal"
		} else if imc > 25.8 {
			mensagem = "acima do peso ideal"
		}
	} else {
		fmt.Println("Sexo inválido")
		return
	}

	fmt.Printf("O IMC da pessoa é %.2f e ela está %s\n", imc, mensagem)
}
