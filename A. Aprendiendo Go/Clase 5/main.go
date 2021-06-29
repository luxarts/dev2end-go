package main

import "fmt"

func main() {
	// En Go solo existe el bucle for el cual se utiliza para repetir
	// un bloque de código una determinada cantidad de veces

	vueltas := 10

	// El primer término es el inicializador. En este caso crea una variable 'v' con el valor 0.
	// El segundo término es la condición. En este caso la condición es que de 'v' sea menor que el valor 'vueltas'.
	// El tercer término es la acción si se cumple la condición. En este caso se incrementa 1 el valor de 'v'.
	for v := 0; v < vueltas; v++ {
		fmt.Println("Vuelta: ", v)
	}

	// Se pueden omitir el bloque de inicialización y de acción, dejando solo el de condición.
	// De esta manera, el bucle for se comporta como un bucle while, el cual se utiliza para repetir
	// un bloque de código siempre que la condición sea verdadera

	condicion := true

	contador := 0

	// While
	for condicion == true {
		if contador == 10 { // Si el contador llega a 10, la condición se hace false
			condicion = false
		} else {
			contador++
		}
	}

	// En Go existe la palabra clave 'range' la cual se utiliza para obtener el índice y el valor de los valores
	// de vector o mapa de manera secuencial.

	frutas := []string{"banana", "manzana", "pera"}

	for indice, nombre := range frutas {
		fmt.Println(indice, nombre)
	}
}
