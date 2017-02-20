package main

import (
	"fmt"
)

func main() {
	var miCafe = Cafe{"espresso", 5.22, false, 0}

	fmt.Print(miCafe)

}

type Cafe struct {
	nombre string
	precio float64
	azucar bool
	leche  int
}
