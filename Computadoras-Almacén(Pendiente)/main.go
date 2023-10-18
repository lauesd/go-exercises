package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	
)

// Estructura para representar una computadora
type Computer struct {
	Modelo          string   `json:"modelo"`
	Marca           string   `json:"marca"`
	Precio          float64  `json:"precio"`
	Especificaciones []string `json:"especificaciones"`
}

// Función para guardar las computadoras en el archivo JSON
func guardarComputadoras(computadoras []Computer) error {
	// Convertir las computadoras a formato JSON
	computadorasJSON, err := json.Marshal(computadoras)
	if err != nil {
		return err
	}

	// Escribir las computadoras en el archivo JSON
	err = ioutil.WriteFile("computadoras.json", computadorasJSON, 0644)
	if err != nil {
		return err
	}

	return nil
}

// Función para obtener todas las computadoras del archivo JSON
func obtenerComputadoras() ([]Computer, error) {
	// Leer el contenido del archivo JSON
	fileContent, err := ioutil.ReadFile("computadoras.json")
	if err != nil {
		return nil, err
	}

	// Crear una lista para almacenar las computadoras
	var computadoras []Computer

	// Deserializar el contenido JSON en la lista de computadoras
	err = json.Unmarshal(fileContent, &computadoras)
	if err != nil {
		return nil, err
	}

	return computadoras, nil
}

func main() {
	// Crear una lista de computadoras de ejemplo
	computadoras := []Computer{
		{
			Modelo:          "MacBook Pro",
			Marca:           "Apple",
			Precio:          1999.99,
			Especificaciones: []string{"Procesador Intel Core i7", "16 GB RAM", "512 GB SSD"},
		},
		{
			Modelo:          "ThinkPad X1 Carbon",
			Marca:           "Lenovo",
			Precio:          1599.99,
			Especificaciones: []string{"Procesador Intel Core i5", "8 GB RAM", "256 GB SSD"},
		},
	}

	// Guardar las computadoras en el archivo JSON
	err := guardarComputadoras(computadoras)
	if err != nil {
		fmt.Println("Error al guardar las computadoras:", err)
		return
	}

	// Obtener todas las computadoras del archivo JSON
	computadorasObtenidas, err := obtenerComputadoras()
	if err != nil {
		fmt.Println("Error al obtener las computadoras:", err)
		return
	}

	// Mostrar la información de las computadoras
	fmt.Println("Computadoras disponibles:")
	for _, computadora := range computadorasObtenidas {
		fmt.Println("Modelo:", computadora.Modelo)
		fmt.Println("Marca:", computadora.Marca)
		fmt.Println("Precio:", computadora.Precio)
		fmt.Println("Especificaciones:", computadora.Especificaciones)
		fmt.Println()
	}
}

