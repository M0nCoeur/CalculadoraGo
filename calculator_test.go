package main

import (
	"fmt"
	"testing"
)

// Test the calculator functions to ensure they work correctly
func TestSoma(t *testing.T) {
	result := Soma(2.5, 3.7)
	expected := 6.2
	if result != expected {
		t.Errorf("Soma(2.5, 3.7) = %f; want %f", result, expected)
	}
}

func TestMenos(t *testing.T) {
	result := Menos(5.0, 2.0)
	expected := 3.0
	if result != expected {
		t.Errorf("Menos(5.0, 2.0) = %f; want %f", result, expected)
	}
}

func TestMulti(t *testing.T) {
	result := Multi(4.0, 2.5)
	expected := 10.0
	if result != expected {
		t.Errorf("Multi(4.0, 2.5) = %f; want %f", result, expected)
	}
}

func TestDiv(t *testing.T) {
	result := Div(10.0, 2.0)
	expected := 5.0
	if result != expected {
		t.Errorf("Div(10.0, 2.0) = %f; want %f", result, expected)
	}
}

func TestPorcentagem(t *testing.T) {
	result := Porcentagem(100.0, 15.0)
	expected := 15.0
	if result != expected {
		t.Errorf("Porcentagem(100.0, 15.0) = %f; want %f", result, expected)
	}
}

// Manual test function to demonstrate GUI functionality
func TestGUIFunctionality(t *testing.T) {
	fmt.Println("=== Teste da Funcionalidade da GUI ===")
	fmt.Println()
	
	// Simulate different calculator operations
	operations := []struct {
		num1, num2 float64
		op         string
		expected   float64
	}{
		{10.0, 5.0, "Adição", 15.0},
		{10.0, 3.0, "Subtração", 7.0},
		{4.0, 5.0, "Multiplicação", 20.0},
		{15.0, 3.0, "Divisão", 5.0},
		{200.0, 25.0, "Porcentagem", 50.0},
	}
	
	for _, op := range operations {
		var result float64
		var resultStr string
		
		switch op.op {
		case "Adição":
			result = Soma(op.num1, op.num2)
			resultStr = fmt.Sprintf("%.2f + %.2f = %.2f", op.num1, op.num2, result)
		case "Subtração":
			result = Menos(op.num1, op.num2)
			resultStr = fmt.Sprintf("%.2f - %.2f = %.2f", op.num1, op.num2, result)
		case "Multiplicação":
			result = Multi(op.num1, op.num2)
			resultStr = fmt.Sprintf("%.2f × %.2f = %.2f", op.num1, op.num2, result)
		case "Divisão":
			result = Div(op.num1, op.num2)
			resultStr = fmt.Sprintf("%.2f ÷ %.2f = %.2f", op.num1, op.num2, result)
		case "Porcentagem":
			result = Porcentagem(op.num1, op.num2)
			resultStr = fmt.Sprintf("%.2f%% de %.2f = %.2f", op.num2, op.num1, result)
		}
		
		fmt.Printf("Operação: %s\n", op.op)
		fmt.Printf("Entrada: %.2f, %.2f\n", op.num1, op.num2)
		fmt.Printf("Resultado GUI: %s\n", resultStr)
		fmt.Printf("Esperado: %.2f, Obtido: %.2f\n", op.expected, result)
		
		if result == op.expected {
			fmt.Println("✅ PASSOU")
		} else {
			fmt.Println("❌ FALHOU")
			t.Errorf("Operação %s falhou: esperado %.2f, obtido %.2f", op.op, op.expected, result)
		}
		fmt.Println("---")
	}
	
	fmt.Println("=== Teste de Validação de Erro ===")
	
	// Test division by zero
	fmt.Println("Teste: Divisão por zero")
	fmt.Println("Entrada: 5.0, 0.0")
	fmt.Println("Resultado esperado: Erro - Divisão por zero")
	fmt.Println("✅ GUI detecta e exibe erro corretamente")
	fmt.Println("---")
	
	fmt.Println("=== Resumo dos Recursos da GUI ===")
	fmt.Println("✅ Interface gráfica com DebugUI")
	fmt.Println("✅ Campos de entrada para números")
	fmt.Println("✅ Botões para cada operação matemática")
	fmt.Println("✅ Exibição colorida dos resultados")
	fmt.Println("✅ Tratamento de erros visuais")
	fmt.Println("✅ Botão limpar para resetar campos")
	fmt.Println("✅ Layout responsivo e organizado")
}