package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

// Las funciones de test deben empezar con la palabra Test y deben recibir un puntero de la estructura testing.T
func TestDividir_OK(t *testing.T){
	// Una forma de escribir tests legibles es usando el lenguaje Gherkin, el cual consiste en dividir el test en 3
	// partes: Given (dado), When (cuando), Then (entonces)

	// Given (dadas estas variables e inicializaciones)
	var a float32 = 10
	var b float32 = 2

	// When (cuando se ejecute el siguiente código)
	result, err := dividir(a, b)

	// Then (entonces lo siguiente tiene que cumplirse)
	require.Nil(t, err) // El error debe ser nil
	require.Equal(t, float32(5), result) // El resultado debe ser 5
}

func TestDividir_ErrorAlDividirPorCero(t *testing.T){
	// Given
	var a float32 = 10
	var b float32 = 0

	// When
	result, err := dividir(a, b)

	// Then
	require.NotNil(t, err) // El error debe ser distinto de nil
	require.Empty(t, result) // El resultado debe estar vacío (0)
	require.Equal(t, "no se puede dividir por 0", err.Error()) // El mensaje del error debe ser "no se puede dividir por 0"
}