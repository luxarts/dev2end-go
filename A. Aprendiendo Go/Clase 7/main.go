package main

import "fmt"

// Una estructura se usa para definir un tipo de datos que agrupa variables (propiedades)
type automovil struct {
	marca string
	ruedas int
	motor motor // Una estructura puede tener una propiedad que sea de tipo estructura también
}

type motor struct {
	combustible string
	acelerando bool
}

// Se pueden asignar funciones para una estructura determinada.
// Para ello, se debe indicar la estructura en el receptor de la función
func (a automovil) acelerar() {
	// La variable 'a' representa a un automovil y contiene las propiedades del mismo
	a.motor.acelerando = true

	fmt.Println("Pisando el pedal acelerador...")
}
func (a automovil) desacelerar() {
	a.motor.acelerando = false

	fmt.Println("Soltando el pedal acelerador...")
}

func newAutomovil(marca string, combustible string) automovil {
	return automovil{
		marca:  marca,
		ruedas: 4,
		motor:  motor{
			combustible: combustible,
		},
	}
}

func main() {
	// Para crear una variable con la estructura de automovil se puede hacer de la siguiente manera
	auto1 := automovil{
		marca:  "Volkswagen",
		ruedas: 4,
		motor:  motor{
			combustible: "GNC",
		},
	}
	// Nota: Si a una propiedad no se le asigna valor, esta tomará el "zero value" correspondiente. En este caso la
	// propiedad 'acelerando' dentro de 'motor' tendrá como valor false.

	// Sin embargo es una buena práctica usar una función constructora que devuelva una estructura
	auto2 := newAutomovil("Volkswagen", "GNC")

	// Para usar las funciones de la estructura lo hacemos de la siguiente manera
	auto1.acelerar()
	auto2.desacelerar()
}