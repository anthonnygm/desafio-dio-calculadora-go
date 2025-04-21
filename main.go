package main

import (
	"bufio"
	"desafio-dio-calculadora-go/expressions"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("🔢 Bem vindo(a) à Calculadora!")

	scanner := bufio.NewScanner(os.Stdin)
	var result float64
	var useResult bool

	for {
		if !useResult {
			fmt.Print("Digite o número inicial: ")
			scanner.Scan()
			numStr := scanner.Text()
			num, err := strconv.ParseFloat(numStr, 64)
			if err != nil {
				fmt.Println("❌ Valor inválido. Tente novamente.")
				continue
			}
			result = num
		}

		fmt.Println("\nEscolha a operação:")
		fmt.Println("1. Soma (+)")
		fmt.Println("2. Subtração (-)")
		fmt.Println("3. Multiplicação (*)")
		fmt.Println("4. Divisão (/)")
		fmt.Print("Digite o número da operação: ")

		scanner.Scan()
		operacao := scanner.Text()

		fmt.Print("Digite o próximo número: ")
		scanner.Scan()
		nextStr := scanner.Text()
		next, err := strconv.ParseFloat(nextStr, 64)
		if err != nil {
			fmt.Println("❌ Valor inválido. Tente novamente.")
			continue
		}

		switch operacao {
		case "1":
			result = expressions.Sum(result, next)
		case "2":
			result = expressions.Subtract(result, next)
		case "3":
			result = expressions.Multiply(result, next)
		case "4":
			res, err := expressions.Divide(result, next)
			if err != nil {
				fmt.Println("❌ Erro: ", err)
				continue
			}
			result = res
		default:
			fmt.Println("❌ Operação inválida.")
			continue
		}

		fmt.Printf("✅ Resultado %s\n", FormatResult(result))

		fmt.Println("\nO que deseja fazer agora?")
		fmt.Println("1. Continuar calculando com o resultado")
		fmt.Println("2. Começar nova operação")
		fmt.Println("3. Sair")

		scanner.Scan()
		opcao := strings.TrimSpace(scanner.Text())

		if opcao == "1" {
			useResult = true
		} else if opcao == "2" {
			useResult = false
		} else if opcao == "3" {
			fmt.Println("👋 Saindo da calculadora. Até a próxima!")
			break
		} else {
			fmt.Println("❌ Opção inválida. Encerrando.")
			break
		}
	}
}

func FormatResult(result float64) string {
	if result == float64(int64(result)) {
		return fmt.Sprintf("%d", int64(result))
	}
	return fmt.Sprintf("%.2f", result)
}
