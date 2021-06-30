package main

// Importamos el código de los paquetes internos
import (
	"clase9/vehiculo"
)

// En Go los proyectos se llaman módulos.
// El archivo go.mod contiene la versión de Go y las dependencias (librerías) que se necesitan en el proyecto.
// Este archivo se genera con el comando: go mod init nombre_del_proyecto

func main(){
	// Si una función, variable o constante está en otro paquete, para usarse se debe poner el nombre del paquete
	// primero seguido de un punto
	auto := vehiculo.NewVehiculo("automovil", "Volkswagen", "GNC")
	auto.Acelerar()
	auto.Desacelerar()

	moto := vehiculo.NewVehiculo("motocicleta", "Honda", "Super")
	moto.Acelerar()
	moto.Desacelerar()
}
