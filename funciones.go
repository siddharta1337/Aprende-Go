package main

import (
	"fmt"
)

func main() {

	fmt.Print(ganado(1))
}

func ganado(numero int) (int, string) {

	vacas := func() int {
		return numero * 10
	}

	return vacas() + 1, " vacas"
}
