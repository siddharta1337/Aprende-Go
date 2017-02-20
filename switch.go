package main

import (
	"fmt"
	"time"
)

func main() {

	tiempo := time.Now()

	switch diaDeHoy := tiempo.Weekday(); diaDeHoy {

	case 0:
		fmt.Println("Descanso de Domingo")
	case 1:
		fmt.Println("Odio los Lunes")
	default:
		fmt.Println("Trabajo")

	}

}
