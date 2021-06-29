package main

import "fmt"

// Como vimos en la clase anterior, las estructuras pueden propiedades y funciones asignadas
// Las interfaces se utilizan para definir un tipo de dato que agrupa funciones.
// Cuando una estructura contiene esas funciones se dice que implementa a esa interfaz

// Definimos la estructura automovil
type automovil struct {
	marca string
	ruedas int
	motor motor
}
type motor struct {
	combustible string
	acelerando bool
}

// Definimos las funciones que tendrá un automovil y lo que harán
func (a automovil) acelerar() {
	a.motor.acelerando = true
	fmt.Println("Pisando el pedal acelerador...")
}
func (a automovil) desacelerar() {
	a.motor.acelerando = false
	fmt.Println("Soltando el pedal acelerador...")
}

// Definimos el constructor de un automovil
func newAutomovil(marca string, combustible string) automovil {
	return automovil{
		marca:  marca,
		ruedas: 4,
		motor:  motor{
			combustible: combustible,
		},
	}
}

// Definimos la estructura motocicleta
type motocicleta struct {
	marca string
	ruedas int
	motor motor
}

// Definimos las funciones que tendrá una motocicleta y lo que harán
func (m motocicleta) acelerar() {
	m.motor.acelerando = true
	fmt.Println("Girando el acelerador hacia abajo...")
}
func (m motocicleta) desacelerar() {
	m.motor.acelerando = false
	fmt.Println("Girando el acelerador hacia arriba...")
}

// Definimos el constructor de una motocicleta
func newMotocicleta(marca string, combustible string) motocicleta {
	return motocicleta{
		marca:  marca,
		ruedas: 2,
		motor:  motor{
			combustible: combustible,
		},
	}
}

// Definimos una interfaz vehículo
type vehiculo interface {
	acelerar()
	desacelerar()
}

// Definimos el constructor de un vehiculo
func newVehiculo(tipo string, marca string, combustible string) vehiculo {
	// Creamos una variable de tipo vehiculo. Como no le asignamos valor, este será 'nil'
	var v vehiculo

	if tipo == "motocicleta" {
		v = newMotocicleta(marca, combustible)
	} else if tipo == "automovil" {
		v = newAutomovil(marca, combustible)
	}

	// Como 'motocicleta' y 'automovil' tienen las misma funciones que la interfaz, decimos que ambas son
	// implementaciones

	return v
}

func main() {
	// Ahora para crear un automovil, en vez de hacerlo con el constructor del auto directamente, podemos hacerlo con el
	// constructor de la interfaz vehiculo
	auto := newVehiculo("automovil", "Volkswagen", "GNC")
	auto.acelerar()
	auto.desacelerar()

	// Para crear una motocicleta, simplemente cambiamos el valor que le pasamos al constructorç
	moto := newVehiculo("motocicleta", "Honda", "Super")
	moto.acelerar()
	moto.desacelerar()
}