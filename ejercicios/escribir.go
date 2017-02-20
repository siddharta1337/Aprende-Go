package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	contenido := []byte("Esto es un dato almacenado desde GO")

	datos := ioutil.WriteFile("archivo.txt", contenido, 0755)

	mostrarError(datos)

	fmt.Print("info lista")

}

func mostrarError(e error) {
	if e != nil {
		panic(e)
	}
}
