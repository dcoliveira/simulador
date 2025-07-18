# Simulador Rover de Marte
Este projeto simula o movimento de rovers robóticos em um platô em Marte, seguindo instruções.

## 🚀 Tecnologias utilizadas

- Golang 1.21.3 (Não foi testado em versões anteriores)

---

## ⚙️ Requisitos

- Golang 1.21.3

---

# Estrutura do Projeto
O código é organizado em arquivos Go no mesmo diretório:

main.go: Lógica principal e ponto de entrada.

rover.go: Define a estrutura Rover e seus métodos de movimento/rotação.

mock_data.go: Contém a entrada de dados mocada.

main_test.go: Inclui testes unitários para o projeto.

# Como Rodar o Projeto

1. Rodar com Entrada Mocada (para demonstração)

Linux/macOS: 

```bash
go run .
```

# Como Rodar os Testes
No diretório do projeto, execute:

```bash
go test
```
