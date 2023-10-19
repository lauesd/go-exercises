package main

import (
	"fmt"
	"strings"
)

func main() {
	var respuesta1 string
	var respuesta2 string

	fmt.Print("¿Te gustan los videojuegos?: ")
	fmt.Scanln(&respuesta1)

	fmt.Print("¿Te gustarìa meterte a nuestro club?: ")
	fmt.Scanln(&respuesta2)

	respuesta1 = strings.ToLower(respuesta1) 
	respuesta2 = strings.ToLower(respuesta2)

	if respuesta1 == "si" && respuesta2 == "si"{
		fmt.Println("Puedes entrar a esta zona.")
	} else {
		fmt.Println("No puedes entrar a esta zona")
	} 
}