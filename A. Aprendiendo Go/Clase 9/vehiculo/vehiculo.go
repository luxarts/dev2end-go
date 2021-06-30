package vehiculo

// Como la estructura motor la usan todos los vehiculos, es mejor definirla en este archivo.
type motor struct {
	combustible string
	acelerando bool
}

// Definimos una interfaz vehículo
type Vehiculo interface {
	Acelerar()
	Desacelerar()
}

// Definimos el constructor de un vehiculo
func NewVehiculo(tipo string, marca string, combustible string) Vehiculo {
	var v Vehiculo

	if tipo == "motocicleta" {
		v = NewMotocicleta(marca, combustible)
	} else if tipo == "automovil" {
		v = NewAutomovil(marca, combustible)
	}

	return v
}

// Nota: Para poder usar las funciones, variables o constantes declaradas desde otro paquete, estas deben comenzar
// con la primera letra en mayúscula