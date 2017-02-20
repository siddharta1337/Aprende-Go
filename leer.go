package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	datos, errorDeLectura := ioutil.ReadFile("archivo.txt")

	mostrarError(errorDeLectura)

	fmt.Println(string(datos))

}

func mostrarError(e error) {
	if e != nil {
		panic(e)
	}
}
