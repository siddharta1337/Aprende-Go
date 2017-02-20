package main

import (
	"fmt"
)

func main() {

	puntos := 100

	if puntos < 100 {
		fmt.Print("puntos incorrectos")
	} else if puntos == 100 {
		fmt.Print("puntos CORRECTOS")
	} else {
		fmt.Print("tus puntos son ", puntos)
	}

}
