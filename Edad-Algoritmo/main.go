package main

import (
	"fmt"
	"strings"
)

func main() {
	var edad int
    var respuesta string

	fmt.Print("¿Podrías ingresar tu edad?: ")
	fmt.Scanln(&edad)

    fmt.Print("¿Traes tu cedula?: ")
	fmt.Scanln(&respuesta)

	respuesta = strings.ToLower(respuesta)
	
	if edad >= 18 && respuesta == "si"{
		fmt.Println("Puedes entrar a esta zona.")
	} else {
		fmt.Println("No puedes entrar a esta zona")
	}

}
