package main

import (
	"fmt"
)

func main() {
	var frutas = []string{"Manzanas", "Uvas", "Bananas"}

	frutas = append(frutas, "Peras")

	fmt.Println(frutas[0:4])
	fmt.Println(len(frutas))
}
