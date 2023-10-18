package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Ingrese un texto: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	texto := scanner.Text()

	numPalabras := contarPalabras(texto)
	
	fmt.Println("El número de palabras es:", numPalabras)
}

func contarPalabras(texto string) int {
	palabras := strings.Split(texto, " ")
	return len(palabras)
}