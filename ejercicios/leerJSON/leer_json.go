package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Pais struct
type Pais struct {
	NOMBRE    string
	POBLACION float32
}

// Paises struct
type Paises struct {
	Paises []Pais `json:"paises"`
}

func main() {

	data, err := os.Open("lista.json")
	if err != nil {
		fmt.Print(err)
	}

	byteValue, _ := ioutil.ReadAll(data)

	var paises Paises

	err = json.Unmarshal(byteValue, &paises)
	if err != nil {
		fmt.Println("error:", err)
	}

	for i := 0; i < len(paises.Paises); i++ {
		println("* " + paises.Paises[i].NOMBRE)
	}
}
