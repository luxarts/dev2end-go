package main

import (
	"fmt"
	"time"
)

// Esta función busca todos los números del 1 al número recibido que sean divisores y los va sumando
func contarDivisores(numero uint64, resultCh chan string) {
	divisores := 1

	for n := uint64(1); n<numero; n++{
		// Si el número recibido dividido por el número actual tiene como resto 0, entonces es un divisor
		if numero%n == 0 {
			divisores++
		}
	}

	// Se envía el resultado al canal
	resultCh <- fmt.Sprintf("Divisores calculados para %d: %d", numero, divisores)
}

// Esta función busca todos los números del 1 al número recibido que sean divisores pero a diferencia de la anterior,
// esta considera que todos los números son divisores, por lo que va restando los que no lo sean
func contarDivisoresInversa(numero uint64, resultCh chan string){
	divisores := numero

	for n := uint64(1); n<numero; n++{
		// Si el número recibido dividido por el número actual tiene resto distinto de 0, entonces no es un divisor
		if numero%n != 0 {
			divisores--
		}
	}

	// Se envía el resultado al canal
	resultCh <- fmt.Sprintf("Divisores calculados de manera inversa para %d: %d", numero, divisores)
}

func main(){
	// Guardamos el tiempo actual
	t := time.Now()

	// Creamos un canal donde se van a recibir los datos de las go rutinas
	resultChan := make(chan string)

	// Ejecutamos ambas funciones en una go rutina
	go contarDivisores(900000000, resultChan)
	go contarDivisoresInversa(900000000, resultChan)

	fmt.Println("Tiempo transcurrido despues de las goroutines: ", time.Since(t))

	// Para leer los valores de una canal se utiliza el operador <- seguido del canal
	resultado := <-resultChan
	// Nota: Si el canal está vacío el código queda frenado hasta recibir algo en el canal que se pueda leer. De
	// esta manera se puede volver a sincronizar el código.
	fmt.Println(resultado)
	fmt.Println("Tiempo transcurrido despues del primer resultado: ", time.Since(t))

	// Podemos leer directamente el canal sin tener que guardarlo en una variable antes
	fmt.Println(<-resultChan)
	fmt.Println("Tiempo transcurrido despues del segundo resultado: ", time.Since(t))


}
