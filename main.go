package main

import (
	"fmt"
	"image/color"
	"strconv"

	"github.com/AllenDang/giu"
)

var (
	num1Input       string = "0"
	num2Input       string = "0"
	resultado       string = ""
	operacao        string = "Selecione uma operação"
	mostrarResultado bool   = false
	erro            string = ""
)

func main() {
	wnd := giu.NewMasterWindow("Calculadora Go", 500, 400, giu.MasterWindowFlagsNotResizable)
	wnd.Run(loop)
}

func loop() {
	giu.SingleWindow().Layout(
		giu.Align(giu.AlignCenter).To(
			giu.Label("🧮 CALCULADORA GO"),
		),
		giu.Separator(),
		giu.Spacing(),

		// Inputs dos números
		giu.Row(
			giu.Label("Primeiro número:"),
			giu.InputText(&num1Input).Size(120),
		),
		giu.Spacing(),
		giu.Row(
			giu.Label("Segundo número:"),
			giu.InputText(&num2Input).Size(120),
		),
		giu.Spacing(),

		// Seletor de operação
		giu.Label("Operação:"),
		giu.Row(
			giu.Button("Adição").Size(80, 30).OnClick(func() { executarOperacao("Adição") }),
			giu.Button("Subtração").Size(80, 30).OnClick(func() { executarOperacao("Subtração") }),
			giu.Button("Multiplicação").Size(80, 30).OnClick(func() { executarOperacao("Multiplicação") }),
		),
		giu.Row(
			giu.Button("Divisão").Size(80, 30).OnClick(func() { executarOperacao("Divisão") }),
			giu.Button("Porcentagem").Size(80, 30).OnClick(func() { executarOperacao("Porcentagem") }),
			giu.Button("Limpar").Size(80, 30).OnClick(limparTudo),
		),
		giu.Spacing(),

		// Mostrar erro se houver
		giu.Condition(erro != "", giu.Layout{
			giu.Style().SetColor(giu.StyleColorText, color.RGBA{255, 76, 76, 255}).To(
				giu.Label("⚠️ " + erro),
			),
			giu.Spacing(),
		}, nil),

		// Mostrar resultado
		giu.Condition(mostrarResultado && resultado != "", giu.Layout{
			giu.Separator(),
			giu.Align(giu.AlignCenter).To(
				giu.Label("Resultado:"),
			),
			giu.Align(giu.AlignCenter).To(
				giu.Style().SetColor(giu.StyleColorText, color.RGBA{76, 255, 76, 255}).To(
					giu.Label(resultado),
				),
			),
		}, nil),

		giu.Spacing(),
		giu.Separator(),
		giu.Align(giu.AlignCenter).To(
			giu.Label("Desenvolvido em Go com Dear ImGui"),
		),
	)
}

func executarOperacao(op string) {
	// Limpar erro anterior
	erro = ""

	// Converter inputs para float64
	num1, err1 := strconv.ParseFloat(num1Input, 64)
	num2, err2 := strconv.ParseFloat(num2Input, 64)

	if err1 != nil {
		erro = "Primeiro número inválido"
		mostrarResultado = false
		return
	}

	if err2 != nil {
		erro = "Segundo número inválido"
		mostrarResultado = false
		return
	}

	// Verificar divisão por zero
	if op == "Divisão" && (num2 == 0 || num1 == 0) {
		erro = "Erro! Divisão por zero."
		mostrarResultado = false
		return
	}

	// Executar operação
	var res float64
	switch op {
	case "Adição":
		res = Soma(num1, num2)
		resultado = fmt.Sprintf("%.2f + %.2f = %.2f", num1, num2, res)
	case "Subtração":
		res = Menos(num1, num2)
		resultado = fmt.Sprintf("%.2f - %.2f = %.2f", num1, num2, res)
	case "Multiplicação":
		res = Multi(num1, num2)
		resultado = fmt.Sprintf("%.2f × %.2f = %.2f", num1, num2, res)
	case "Divisão":
		res = Div(num1, num2)
		resultado = fmt.Sprintf("%.2f ÷ %.2f = %.2f", num1, num2, res)
	case "Porcentagem":
		res = Porcentagem(num1, num2)
		resultado = fmt.Sprintf("%.2f%% de %.2f = %.2f", num2, num1, res)
	default:
		erro = "Operação inválida"
		mostrarResultado = false
		return
	}

	operacao = op
	mostrarResultado = true
}

func limparTudo() {
	num1Input = "0"
	num2Input = "0"
	resultado = ""
	operacao = "Selecione uma operação"
	mostrarResultado = false
	erro = ""
}
