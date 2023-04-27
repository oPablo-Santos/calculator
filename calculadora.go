package main

import (
    "fmt"
)

func main() {
    var operacao string
    var num1, num2 float64

    fmt.Println("Digite a operação (+, -, *, /):")
    fmt.Scanln(&operacao)

    fmt.Println("Digite o primeiro número:")
    fmt.Scanln(&num1)

    fmt.Println("Digite o segundo número:")
    fmt.Scanln(&num2)

    switch operacao {
    case "+":
        resultado := num1 + num2
        fmt.Printf("O resultado da soma é %.2f\n", resultado)
    case "-":
        resultado := num1 - num2
        fmt.Printf("O resultado da subtração é %.2f\n", resultado)
    case "*":
        resultado := num1 * num2
        fmt.Printf("O resultado da multiplicação é %.2f\n", resultado)
    case "/":
        if num2 == 0 {
            fmt.Println("Erro: não é possível dividir por zero")
        } else {
            resultado := num1 / num2
            fmt.Printf("O resultado da divisão é %.2f\n", resultado)
        }
    default:
        fmt.Println("Operação inválida!")
    }
}
