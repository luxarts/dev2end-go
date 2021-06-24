package main

import (
	"errors"
	"fmt"
)

func main() {
	// Hay dos formas de declarar una variable
	var nombre string = "Bruce" // var nombreDeLaVariable tipo = valor

	// o
	apellido := "Wayne" // nombreDeLaVariable := valor

	// En este segunda forma, el tipo de datos se determina por el valor

	fmt.Println(nombre, apellido)

	// Se pueden crear multiples variables en una sola linea
	x, y := 1, 2
	fmt.Println(x, y)

	/* -------------------------------------------------------------------- */

	// Para declarar una constante se realiza de la siguiente manera
	const pi = 3.14 // En el caso de las constantes no se necesita especificar el tipo de datos

	/* -------------------------------------------------------------------- */

	// Los tipos de datos que existen en Go son
	var numero int = -2147483647                             // La cantidad de bits depende del CPU, cómo minimo es un int32
	var numeroSinSigno uint = 0                              // La cantidad de bits depende del CPU, cómo minimo es un uint32
	var numeroDe8bits int8 = -127                            // -127 a +127
	var numeroDe8bitsSinSigno uint8 = 255                    // 0 a 255
	var numeroDe16bits int16 = -32767                        // -32767 a 32767
	var numeroDe16bitsSinSigno uint16 = 65535                // 0 a 65535
	var numeroDe32bits int32 = -2147483647                   // -2147483648 a 2147483648
	var numeroDe32bitsSinSigno uint32 = 4294967295           // 0 a 4294967295
	var numeroDe64bits int64 = -9223372036854775808          // -9223372036854775808 a 9223372036854775807
	var numeroDe64bitsSinSigno uint64 = 18446744073709551615 // 0 a 18446744073709551615
	var numeroDecimalDe32bits float32 = 1.23456789
	var numeroDecimalDe64bits float64 = 1.23456789
	var numeroComplejoDe64bits complex64 = 1.234 + 123.45i   // Real float32, imaginario float32
	var numeroComplejoDe128bits complex128 = 1.234 + 123.45i // Real float64, imaginario float64

	var cadenaDeCaracteres string = "texto"

	var unError error = errors.New("Esto es un error")

	var booleano bool = true

	// Los punteros almacenan la dirección de memoria en vez del valor
	var punteroDeUnInt *int = nil // El valor 'nil' es un valor nulo que se puede interpretar como "no existe"

	// Para obtener la dirección de memoria en vez del valor se utiliza el símbolo '&' antes de la variable
	var punteroDeUnString *string = &cadenaDeCaracteres

	fmt.Printf("Numero: %v\n", numero)
	fmt.Printf("Numero s/ signo: %v\n", numeroSinSigno)
	fmt.Printf("Numero 8 bits: %v\n", numeroDe8bits)
	fmt.Printf("Numero 8 bits s/ signo: %v\n", numeroDe8bitsSinSigno)
	fmt.Printf("Numero 16 bits: %v\n", numeroDe16bits)
	fmt.Printf("Numero 16 bits s/ signo: %v\n", numeroDe16bitsSinSigno)
	fmt.Printf("Numero 32 bits: %v\n", numeroDe32bits)
	fmt.Printf("Numero 32 bits s/ signo: %v\n", numeroDe32bitsSinSigno)
	fmt.Printf("Numero 64 bits: %v\n", numeroDe64bits)
	fmt.Printf("Numero 64 bits s/ signo: %v\n", numeroDe64bitsSinSigno)
	fmt.Printf("Numero decimal de 32 bits: %v\n", numeroDecimalDe32bits)
	fmt.Printf("Numero decimal de 64 bits: %v\n", numeroDecimalDe64bits)
	fmt.Printf("Numero complejo de 64 bits: %v\n", numeroComplejoDe64bits)
	fmt.Printf("Numero complejo de 128 bits: %v\n", numeroComplejoDe128bits)
	fmt.Printf("Cadena de caracteres: %v\n", cadenaDeCaracteres)
	fmt.Printf("Error: %v\n", unError)
	fmt.Printf("Booleano: %v\n", booleano)
	fmt.Printf("Puntero de un int (vacio): %v\n", punteroDeUnInt)
	fmt.Printf("Puntero de string: %v\n", punteroDeUnString)
}
