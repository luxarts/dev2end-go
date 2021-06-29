package main

import (
	"fmt"
	"time"
)

// Go posee una forma de ejecutar dos funciones de manera concurrente. Esto quiere decir que en vez de espera que una
// función termine para ejecutar la siguiente, se ejecutan ambas de manera alternada hasta completarse.

func main() {
	// Guardamos el tiempo actual
	tiempo := time.Now()

	saludar("Juan")
	saludar("Pedro")

	fmt.Println("Tiempo transcurrido de manera secuencial (sincrónica):",time.Since(tiempo))

	// Esto como vemos tarda 2 segundos ya que despues de cada saludo espera 1 segundo

	// Guardamos el tiempo actual
	tiempo = time.Now()

	// Usamos la palabra 'go' antes de la función para indicar que la función se ejecutará de manera concurrente
	go saludar("Juan")
	go saludar("Pedro")


	fmt.Println("Tiempo transcurrido con go rutinas (asincrónica):",time.Since(tiempo))

	// Este sleep es necesario para esperar a que las go routinas se ejecuten ya que al ser tán rápida la ejecución el
	// programa terminaría antes de que el saludo sea escrito.
	time.Sleep(1*time.Second)
}

func saludar(nombre string){
	fmt.Println("Hola " + nombre)

	time.Sleep(1*time.Second) // Espera un segundo
}