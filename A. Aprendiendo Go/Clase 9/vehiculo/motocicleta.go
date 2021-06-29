package vehiculo

import "fmt"

// Definimos la estructura motocicleta
type Motocicleta struct {
	marca string
	ruedas int
	motor motor
}

// Definimos las funciones que tendrá una motocicleta y lo que harán
func (m Motocicleta) Acelerar() {
	m.motor.acelerando = true
	fmt.Println("Girando el acelerador hacia abajo...")
}
func (m Motocicleta) Desacelerar() {
	m.motor.acelerando = false
	fmt.Println("Girando el acelerador hacia arriba...")
}

// Definimos el constructor de una motocicleta
func NewMotocicleta(marca string, combustible string) Motocicleta {
	return Motocicleta{
		marca:  marca,
		ruedas: 2,
		motor:  motor{
			combustible: combustible,
		},
	}
}
