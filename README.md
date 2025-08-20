# CalculadoraGo

## Uma calculadora em Go com interface gráfica usando DebugUI

Esta calculadora oferece duas interfaces:
1. **Interface Gráfica (GUI)** - Usando DebugUI da Ebitengine
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
go build -tags console -o calculadora-console
./calculadora-console
```

### Dependências

- Go 1.23+
- DebugUI (através da biblioteca `github.com/ebitengine/debugui`)
- Ebitengine v2 para renderização gráfica
- X11 development libraries (Linux)

### Instalação das Dependências no Ubuntu/Debian

```bash
sudo apt-get update
sudo apt-get install -y libx11-dev libxrandr-dev libxinerama-dev libxcursor-dev libxi-dev libgl1-mesa-dev libxxf86vm-dev
```

### Build

```bash
# Interface gráfica (padrão)
go build -o calculadora-gui

# Interface console  
go build -tags console -o calculadora-console
```

### Testes

```bash
# Executar testes das funções matemáticas
go test -tags test
```

### Tecnologias Utilizadas

- **Linguagem**: Go
- **GUI Framework**: DebugUI (parte da Ebitengine)
- **Game Engine**: Ebitengine v2
- **Renderização**: OpenGL/Vulkan (através da Ebitengine)
