package main

import (
	"errors"
	"fmt"
)

type vehiculo interface {
	doblar(sentido string)
}

func NewVehiculo(tipo string) (vehiculo, error) {
	if tipo == "auto" {
		return &auto{}, nil
	} else if tipo == "moto" {
		return &moto{}, nil
	}
	return nil, errors.New("Tipo invalido")
}

type auto struct {
	cantidadDeRuedas int8
	marca            string
	girandoVolante   bool
}

func (a *auto) doblar(sentido string) {
	a.girandoVolante = true
	fmt.Println("Girando las ruedas delanteras hacia la " + sentido)
}

type moto struct {
	cantidadDeRuedas int8
	marca            string
	girandoManubrio  bool
}

func (m *moto) doblar(sentido string) {
	m.girandoManubrio = true
	fmt.Println("Girando la rueda delantera hacia la " + sentido)
}

func main() {
	bmw := auto{
		cantidadDeRuedas: 4,
		marca:            "BMW",
	}

	bmw.doblar("izquierda")

	yamaha := moto{
		cantidadDeRuedas: 2,
		marca:            "Yamaha",
	}

	yamaha.doblar("derecha")

	v1, err := NewVehiculo("moto")

	if err != nil {
		panic(err.Error())
	}

	v1.doblar("izquierda")
}
