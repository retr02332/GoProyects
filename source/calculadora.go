package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func input(prompt string) string {
	fmt.Printf("%s", prompt)
	operation := bufio.NewScanner(os.Stdin)
	operation.Scan()
	output := operation.Text()
	return output
}

func evaluate(operator string, operator1 int, operator2 int) int {
	switch operator {
	case "+":
		return operator1 + operator2

	case "-":
		return operator1 - operator2

	case "*":
		return operator1 * operator2

	case "/":
		return operator1 / operator2

	default:
		return 0
	}
}

func main() {
	/*
		_ "barra abajo" es un wildcard para cuando una funcion
		devuelve algo pero no lo queremos guardar en memoria,
		ta que si la guardamos en memoria, implicaria usarla
	*/
	operation := input("Ingresa una operacion: ")
	operator := input("\nIngresa el operador: ")

	values := strings.Split(operation, operator)
	operator1, _ := strconv.Atoi(values[0])
	operator2, _ := strconv.Atoi(values[1])

	result := evaluate(operator, operator1, operator2)
	fmt.Println("\nEl resultado es: ", result)

}
