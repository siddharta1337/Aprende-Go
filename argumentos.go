package main

import (
	"fmt"
)

func main() {
	alumnos("Hugo", "Francisco", "Erik")
}

func alumnos(alumno ...string) {

	for _, valor := range alumno {
		fmt.Println(valor)
	}

}
