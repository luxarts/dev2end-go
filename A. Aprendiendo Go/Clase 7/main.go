package main

import "fmt"

// Una estructura se usa para definir un tipo de datos que agrupa variables (propiedades)
type automovil struct {
	marca  string
	ruedas int
	motor  motorStruct // Una estructura puede tener una propiedad que sea de tipo estructura también
}

type motorStruct struct {
	combustible string
	acelerando bool
}

// Se pueden asignar funciones para una estructura determinada.
// Para ello, se debe indicar la estructura en el receptor de la función
func (a automovil) acelerar() {
	// La variable 'a' representa a un automovil y contiene las propiedades del mismo
	a.motor.acelerando = true

	fmt.Printf("Pisando el pedal acelerador del auto %s...\n", a.marca)
}
func (a automovil) desacelerar() {
	a.motor.acelerando = false

	fmt.Printf("Soltando el pedal acelerador del auto %s...\n", a.marca)
}

func newAutomovil(marca string, combustible string) automovil {
	return automovil{
		marca:  marca,
		ruedas: 4,
		motor:  motorStruct{
			combustible: combustible,
		},
	}
}

func main() {
	// Para crear una variable con la estructura de automovil se puede hacer de la siguiente manera
	auto1 := automovil{
		marca:  "Volkswagen",
		ruedas: 4,
		motor:  motorStruct{
			combustible: "GNC",
		},
	}
	// Nota: Si a una propiedad no se le asigna valor, esta tomará el "zero value" correspondiente. En este caso la
	// propiedad 'acelerando' dentro de 'motorStruct' tendrá como valor false.

	// Sin embargo es una buena práctica usar una función constructora que devuelva una estructura
	auto2 := newAutomovil("Audi", "Premium")

	// Para usar las funciones de la estructura lo hacemos de la siguiente manera
	auto1.acelerar()
	auto2.desacelerar()
}