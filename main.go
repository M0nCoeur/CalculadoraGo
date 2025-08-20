//go:build !test && !console

package main

import (
	"fmt"
	"image"
	"image/color"
	"strconv"

	"github.com/ebitengine/debugui"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 500
	screenHeight = 400
)

// Game implements ebiten.Game interface
type Game struct {
	debugUI     debugui.DebugUI
	num1Input   string
	num2Input   string
	resultado   string
	erro        string
	showResult  bool
}

// NewGame creates a new Game instance
func NewGame() *Game {
	return &Game{
		num1Input: "0",
		num2Input: "0",
	}
}

// Update implements ebiten.Game interface
func (g *Game) Update() error {
	_, err := g.debugUI.Update(func(ctx *debugui.Context) error {
		// Main window following official debugui example pattern
		ctx.Window("🧮 Calculadora Go", image.Rect(10, 10, screenWidth-10, screenHeight-10), func(layout debugui.ContainerLayout) {
			// Title header
			ctx.Header("CALCULADORA GO", false, func() {
				ctx.Text("Uma calculadora simples feita com DebugUI")
			})
			
			// Input section
			ctx.Header("Entrada de Números", true, func() {
				ctx.Text("Primeiro número:")
				ctx.TextField(&g.num1Input)
				
				ctx.Text("Segundo número:")
				ctx.TextField(&g.num2Input)
			})
			
			// Operations section
			ctx.Header("Operações", true, func() {
				ctx.SetGridLayout([]int{-1, -1}, nil)
				
				// First row of buttons
				ctx.Button("Adição").On(func() {
					g.executarOperacao("Adição")
				})
				
				ctx.Button("Subtração").On(func() {
					g.executarOperacao("Subtração")
				})
				
				// Second row of buttons
				ctx.Button("Multiplicação").On(func() {
					g.executarOperacao("Multiplicação")
				})
				
				ctx.Button("Divisão").On(func() {
					g.executarOperacao("Divisão")
				})
				
				// Third row
				ctx.Button("Porcentagem").On(func() {
					g.executarOperacao("Porcentagem")
				})
				
				ctx.Button("Limpar").On(func() {
					g.limparTudo()
				})
			})
			
			// Results section
			ctx.Header("Resultado", true, func() {
				// Error display
				if g.erro != "" {
					ctx.Text("⚠️ Erro: " + g.erro)
				}
				
				// Result display
				if g.showResult && g.resultado != "" {
					ctx.Text("✅ " + g.resultado)
				}
				
				if g.erro == "" && !g.showResult {
					ctx.Text("Digite os números e escolha uma operação")
				}
			})
			
			// Footer
			ctx.Header("Informações", false, func() {
				ctx.Text("Desenvolvido em Go com DebugUI (Ebitengine)")
			})
		})
		
		return nil
	})
	
	return err
}

// Draw implements ebiten.Game interface
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{48, 48, 48, 255}) // Dark background
	g.debugUI.Draw(screen)
}

// Layout implements ebiten.Game interface
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) executarOperacao(op string) {
	// Limpar erro anterior
	g.erro = ""

	// Converter inputs para float64
	num1, err1 := strconv.ParseFloat(g.num1Input, 64)
	num2, err2 := strconv.ParseFloat(g.num2Input, 64)

	if err1 != nil {
		g.erro = "Primeiro número inválido"
		g.showResult = false
		return
	}

	if err2 != nil {
		g.erro = "Segundo número inválido"
		g.showResult = false
		return
	}

	// Verificar divisão por zero
	if op == "Divisão" && (num2 == 0 || num1 == 0) {
		g.erro = "Erro! Divisão por zero."
		g.showResult = false
		return
	}

	// Executar operação
	var res float64
	switch op {
	case "Adição":
		res = Soma(num1, num2)
		g.resultado = fmt.Sprintf("%.2f + %.2f = %.2f", num1, num2, res)
	case "Subtração":
		res = Menos(num1, num2)
		g.resultado = fmt.Sprintf("%.2f - %.2f = %.2f", num1, num2, res)
	case "Multiplicação":
		res = Multi(num1, num2)
		g.resultado = fmt.Sprintf("%.2f × %.2f = %.2f", num1, num2, res)
	case "Divisão":
		res = Div(num1, num2)
		g.resultado = fmt.Sprintf("%.2f ÷ %.2f = %.2f", num1, num2, res)
	case "Porcentagem":
		res = Porcentagem(num1, num2)
		g.resultado = fmt.Sprintf("%.2f%% de %.2f = %.2f", num2, num1, res)
	default:
		g.erro = "Operação inválida"
		g.showResult = false
		return
	}

	g.showResult = true
}

func (g *Game) limparTudo() {
	g.num1Input = "0"
	g.num2Input = "0"
	g.resultado = ""
	g.showResult = false
	g.erro = ""
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Calculadora Go - DebugUI")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeDisabled)

	game := NewGame()
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
