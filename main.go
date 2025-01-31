package main

import "fmt"

func main() {
	var sair bool

	for !sair {
		var num1, num2 float64
		var operador string

		fmt.Println("Qual operação quer fazer? ")
		fmt.Print("Operadores possíveis (Adição, Subtração, Divisão, Multiplicação e Porcentagem): ")
		fmt.Scan(&operador)

		if operador == "Porcentagem" {
			fmt.Println("Formula da Porcentagem: Total * Porcentagem / 100")
			fmt.Println("Qual o valor total que deve ser considerado?")
		}

		fmt.Print("Digite o primeiro número: ")
		fmt.Scan(&num1)

		fmt.Print("Digite o segundo número: ")
		fmt.Scan(&num2)

		// Podemos utilizar IF dentro de IF e nesse caso, o IF interno só será executado se o IF externo for verdadeiro.
		if operador == "Divisão" {
			if num2 == 0 || num1 == 0 {
				fmt.Print("Erro! Divisão por zero.")
				return
			}
		}

		// O Switch deve ser usado quando vamos verificar uma variável várias vezes.
		switch operador {
		case "Adição":
			fmt.Println(Soma(num1, num2))
		case "Subtração":
			fmt.Println(Menos(num1, num2))
		case "Multiplicação":
			fmt.Println(Multi(num1, num2))
		case "Divisão":
			fmt.Println(Div(num1, num2))
		case "Porcentagem":
			fmt.Println(fmt.Sprint(num2, "% ", "de ", num1, " é: ", Porcentagem(num1, num2)))
		default:
			fmt.Println("Operador inválido")
		}

		var option string
		fmt.Print("Você deseja realizar outra operação? ")
		fmt.Scan(&option)

		if option == "Não" {
			sair = true
		}
	}
}
