package main


import "fmt"


func main() {
	var n int
    fmt.Print("Ingresa un número entero positivo: ")
	fmt.Scanln(&n)

	sum := 0

	for i := 1; i <= n; i++ {
		if i%2 != 0 {
            sum += i
		}
	}

	fmt.Println("La suma de los número impares hasta", n, "es:", sum)
}