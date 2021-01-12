package main

import "fmt"

func main() {

	/*
		defer hace que dicha instruccion que esta marcada por esta keyword,
		sea ejecutada cuando finalice la funcion (todas las demas tareas esten hechas)

		Si todas las tareas estan marcadas con la keyword 'defer', esto causaria que
		la ejecución de la funcion actue de manera inversa.

		Lo anterior sucede gracias a 'LIFO' o mas conocido como 'Last in First out'
		Quiere decir que la ultima instruccion declarada en la función, sera la
		primera en ejecutarse, meintras que si ninguna instruccion esta marcada como
		'defer', entonces la funcion tomara una ejecucion 'FIFO' o mas conocida como
		'First in First out', quiere decir que la primera instruccion en declararse
		en la funcion, sera la primera en ejecutarse.

		En conclusion, la keyword 'defer' es LIFO.
	*/

	// without defer
	fmt.Println("Hola")
	fmt.Println("Mundo")

	// with defer
	defer fmt.Println("Hola")
	fmt.Println("\nMundo")

}
