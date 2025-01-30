package main

import "fmt"

func main() {

	var num1, num2 float64
	var operador string

	fmt.Print("Digite o primeiro número: ")
	fmt.Scan(&num1)

	fmt.Print("Digite o segundo número: ")
	fmt.Scan(&num2)

	fmt.Print("Coloque o operador: +, -, /, *: ")
	fmt.Scan(&operador)

	if operador == "/" {
		if num2 == 0 || num1 == 0 {
			fmt.Print("Erro! Divisão por zero.")
			return
		}
	}

	switch operador {
	case "+":
		fmt.Print(Soma(num1, num2))
	case "-":
		fmt.Print(Menos(num1, num2))
	case "*":
		fmt.Print(Multi(num1, num2))
	case "/":
		fmt.Print(Div(num1, num2))
	default:
		fmt.Println("Operador inválido")
	}

}
