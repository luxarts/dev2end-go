package vehiculo

import "fmt"

// Definimos la estructura automovil
type Automovil struct {
	marca string
	ruedas int
	motor motor
}

// Definimos las funciones que tendrá un automovil y lo que harán
func (a Automovil) Acelerar() {
	a.motor.acelerando = true
	fmt.Println("Pisando el pedal acelerador...")
}
func (a Automovil) Desacelerar() {
	a.motor.acelerando = false
	fmt.Println("Soltando el pedal acelerador...")
}

// Definimos el constructor de un automovil
func NewAutomovil(marca string, combustible string) Automovil {
	return Automovil{
		marca:  marca,
		ruedas: 4,
		motor:  motor{
			combustible: combustible,
		},
	}
}
