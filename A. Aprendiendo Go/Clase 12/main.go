package main

import (
	"errors"
	"fmt"
	"log"
)

// El testeo es una parte muy importante del desarrollo ya que permite saber de una manera rápida si el código funciona
// como debería o no

// Los tests se escriben un archivo con el mismo nombre que el archivo del que se quieren testear las funciones pero
// con el sufijo _test

// Esta función divide un número por otro y devuelve el resultado o un error
func dividir(a float32, b float32) (float32, error){
	// Si el divisor es 0, la operación no se puede realizar y por lo tanto devolvemos un error
	if b == 0 {
		return 0, errors.New("no se puede dividir por 0")
	}

	return a/b, nil
}

func main() {
	r, err := dividir(10, 2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Resultado: ", r)
}

