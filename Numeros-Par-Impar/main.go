package main

import "fmt"

func main() {
	var numero int
	fmt.Print("¿Podrías ingresar un número?: ")
    fmt.Scanln(&numero)

    if (numero % 2 == 0) {
        fmt.Println("El número", numero, "es par.")
    } else {
        fmt.Println("El número", numero, "es impar.")
    }
}