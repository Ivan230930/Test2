package main
 
import (
    "fmt"
    "os"
    "strconv"
    "strings"
)
 
func romanToInt(roman string) int {
 romanMap := map[rune]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
 result := 0
 prev := 0
 
    for _, r := range roman {
 value := romanMap[r]
 result += value
        if prev < value {
 result -= 2 * prev
        }
 prev = value
    }
 
    return result
}
 
func intToRoman(num int) string {
    if num <= 0 {
        panic("Результат должен быть положительным числом")
    }
 
 romanMap := map[int]string{1000: "M", 900: "CM", 500: "D", 400: "CD", 100: "C", 90: "XC", 50: "L", 40: "XL", 10: "X", 9: "IX", 5: "V", 4: "IV", 1: "I"}
 result := ""
 
    for value, numeral := range romanMap {
        for num >= value {
 result += numeral
 num -= value
        }
    }
 
    return result
}
 
func calculate(operand1, operand2 int, operator string) int {
    switch  operator {
    case "+":
        return operand1 + operand2
    case "-": 
        return operand1 - operand2
    case "*":
        return operand1 * operand2
    case "/":
        return operand1 / operand2
    default:
        panic("Неверная арифметическая операция")
    }
}

func main() {
 args := os.Args[1:]
    if len(args) <3 {
        panic("Неверное количество аргументов")
    }
 
 operand1, err := strconv.Atoi(args[0])
    if err != nil {
 operand1 = romanToInt(args[0])
    }
 operand2, err := strconv.Atoi(args[2])
    if err != nil {
 operand2 = romanToInt(args[2])
    }
 
    if operand1 < 1 || operand1 > 10 || operand2 < 1 || operand2 > 10 {
        panic("Числа должны быть от 1 до 10 включительно")
    }
 
    if strings.ContainsAny(args[0], "IVXLCDM") && strings.ContainsAny(args[2], "IVXLCDM") {
        panic("Нельзя использовать одновременно арабские и римские цифры")
    }
 
 result := calculate(operand1, operand2, args[1])
 
    if strings.ContainsAny(args[0], "IVXLCDM") {
 fmt.Println(intToRoman(result))
    } else {
 fmt.Println(result)
    }
}