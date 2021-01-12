package main

import (
	"fmt"

	"./src/banner"
	"./src/product"
	"./src/stdin"
	"./src/store"
)

func main() {
	fmt.Println(banner.ShowBanner())
	userOption := stdin.Input("¿Que desea comprar hoy?\n\n>>> ")
	producto := product.Product{Name: userOption, Price: 2.500, Flavor: "naranja", Expire: "25/ago/2021"}
	tienda := store.Store{Name: "Tienda pepito", Description: "Calidad precio", Product: producto}
	fmt.Printf("\nHas buscado: %s\n\n", tienda.Product.Name)
}
