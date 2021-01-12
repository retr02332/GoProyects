package main

import "fmt"

func main() {

	// Continue y break son soportados por los for de go (todos los aqui presentes)

	// for de c/c++
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	// for de python
	var number int = 1
	fmt.Println("\n")
	for number <= 10 {
		fmt.Println(number)
		number++
	}

	// for de go
	limit := 10
	number = 0
	fmt.Println("\n")
	for {
		fmt.Println(number)
		if number == limit {
			break
		}
		number++
	}

	// for each de go (el mas elegante de los for)
	fmt.Println("\n")
	for index, number := range []int{1, 2, 3, 4, 5} {
		fmt.Println("Index: ", index, "Number: ", number)
	}
}
