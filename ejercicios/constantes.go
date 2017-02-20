package main

import (
	"fmt"
)

func main() {

	const estatico string = "valor 1"
	dinamico := "valor 2"

	/*
	 estatico = false
	*/
	dinamico = "valor 3"

	fmt.Print(estatico, dinamico)
}
