package main

import "fmt"

// Se crea una función que recibe un string y devuelve un string
func saludar(nombre string) string {
	return "Hola " + nombre // Concatena el string 'Hola' con la variable recibida
}

func main() {
	// Se llama a la función 'saludar' con el parámetro 'Bruce'
	saludo := saludar("Bruce")
	fmt.Println(saludo)

	// Función anónima
	despedida := func(nombre string) string {
		return "Adios " + nombre
	}("Bruce") // El último par de paréntesis es el que llama a la función

	fmt.Println(despedida)
}
