# CalculadoraGo

## Uma calculadora em Go com interface gráfica usando Dear ImGui

Esta calculadora oferece duas interfaces:
1. **Interface Gráfica (GUI)** - Usando Dear ImGui através da biblioteca giu
2. **Interface de Terminal** - Versão console original

### Funcionalidades

- ➕ Adição
- ➖ Subtração  
- ✖️ Multiplicação
- ➗ Divisão
- 📊 Cálculo de Porcentagem

### Como Usar

#### Interface Gráfica (Recomendada)
```bash
go run .
```

A interface gráfica oferece:
- Campos de entrada intuitivos para os números
- Botões para cada operação matemática
- Exibição visual dos resultados
- Tratamento de erros com mensagens coloridas
- Botão para limpar todos os campos

#### Interface de Console (Backup)
```bash
go run main_console.go operator.go
```

### Dependências

- Go 1.23+
- Dear ImGui (através da biblioteca `github.com/AllenDang/giu`)
- OpenGL e GLFW (para renderização gráfica)

### Instalação das Dependências no Ubuntu/Debian

```bash
sudo apt-get update
sudo apt-get install -y libgl1-mesa-dev libglfw3-dev libglew-dev
```

### Build

```bash
# Interface gráfica
go build -o calculadora-gui

# Interface console  
go build -o calculadora-console main_console.go operator.go
```

### Screenshot

![Calculadora Go - Interface Gráfica](screenshot.png)

### Tecnologias Utilizadas

- **Linguagem**: Go
- **GUI Framework**: giu (Go bindings para Dear ImGui)
- **Renderização**: OpenGL
- **Sistema de Janelas**: GLFW
