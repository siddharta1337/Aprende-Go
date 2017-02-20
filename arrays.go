package main

import (
	"fmt"
)

func main() {

	var frutas [2][2]string

	// Fruta 1
	frutas[0][0] = "Manzanas"
	frutas[0][1] = "Golden"

	// Fruta 2
	frutas[1][0] = "Uvas"
	frutas[1][1] = "Moscatel"

	fmt.Print(frutas[1][1])

}
