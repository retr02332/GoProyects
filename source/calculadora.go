package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	option := bufio.NewScanner(os.Stdin)
	option.Scan()
	var operation string = option.Text()
	fmt.Println(operation)
	valores := strings.Split(operation, "+")
	fmt.Println(valores)
	fmt.Println(valores[0] + valores[1])
	operator1, _ := strconv.Atoi(valores[0])
	operator2, _ := strconv.Atoi(valores[1])
	fmt.Println(operator1 + operator2)
}
