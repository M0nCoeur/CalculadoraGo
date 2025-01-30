package main

import "fmt"

func main() {

	var num1, num2 float64
	var operador string

	fmt.Print("Qual operação quer fazer? ")
	fmt.Print("Operadores possíveis (+, -, /, *, porcentagem)")
	fmt.Scan(&operador)

	if	operador == ""

	/* Aaa?

porcentagem

Formula da porcentagem: total * porcentagem / 100

Qual é o valor total que deve considerado?

10

Em qual porcentagem você quer saber do total?

2

2% de 10 é: 0,2.

Gostaria de realizar outra operação?



*/

	fmt.Print("Digite o primeiro número: ")
	fmt.Scan(&num1)

	fmt.Print("Digite o segundo número: ")
	fmt.Scan(&num2)

	fmt.Print("Coloque o operador: +, -, /, *: ")
	fmt.Scan(&operador)

	// Podemos utilizar IF dentro de IF e nesse caso, o IF interno só será executado se o IF externo for verdadeiro.
	if operador == "/" {
		if num2 == 0 || num1 == 0 {
			fmt.Print("Erro! Divisão por zero.")
			return
		}
	}

	// O Switch deve ser usado quando vamos verificar uma variável várias vezes.
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
