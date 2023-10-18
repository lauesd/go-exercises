package main

import (
	"fmt"
	"os"
	"strconv"
)


type Calculator struct {
	result float64
}


func (c *Calculator) add(a, b float64) {
	c.result = a + b
}


func (c *Calculator) subtract(a, b float64) {
	c.result = a - b
}


func (c *Calculator) multiply(a, b float64) {
	c.result = a * b
}


func (c *Calculator) divide(a, b float64) {
	if b != 0 {
		c.result = a / b
	} else {
		fmt.Println("Error: Division by zero")
	}
}


func main() {
	c := Calculator{}

	fmt.Println("Welcome to the Calculator")

	fmt.Print("Enter first operand: ")
	num1, _ := strconv.ParseFloat(readInput(), 64)
	
	fmt.Print("Enter second operand: ")
	num2, _ := strconv.ParseFloat(readInput(), 64)

	fmt.Println("Select operation:")
	fmt.Println("1, Add")
	fmt.Println("2, Substract")
	fmt.Println("3, Multiply")
	fmt.Println("4, Divide")

	fmt.Print("Enter operation number: ")
	operation, _ := strconv.Atoi(readInput())

	switch operation {
	case 1:
		c.add(num1, num2)
	case 2:
		c.subtract(num1, num2)
	case 3:
		c.multiply(num1, num2)
	case 4:
		c.divide(num1, num2)
	default:
		fmt.Println("Invalid operation")
		os.Exit(1)
	}
    fmt.Println("Result:", c.result)
}

func readInput() string {
	var input string
	fmt.Scanln(&input)
	return input
}