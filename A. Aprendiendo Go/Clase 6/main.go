package main

import "fmt"

func main() {
	// Al declarar una variable esta se guarda en un registro de la memoria RAM.
	// Podemos imaginar la memoria RAM como una cajonera, en donde cada registro es un cajón.
	// Los cajones pueden guardar cosas (valores) pero también tienen una posición (dirección).
	// Los punteros son variables que en vez de contener un valor, contienen la dirección de memoria de otra variable

	var a int = 10 // Ejemplo: dirección de a = 0x1234
	var b int = 20 // Ejemplo: dirección de b = 0x5678

	// Nota: Como la memoria RAM tiene muchas posiciones, se usa el sistema hexadecimal para referirse a cada una ya
	// que con el sistema decimal nos quedarían números muy grandes.

	// El símbolo '&' indica que se quiere obtener la dirección de la variable en vez de su valor
	// El símbolo '*' indica que lo que se guardará es una dirección
	var pa *int = &a // Ejemplo: pa = 0x1234
	var pb *int = &b // Ejemplo: pb = 0x5678

	fmt.Printf("a vale %v y su dirección es %v\n", a, pa)
	fmt.Printf("b vale %v y su dirección es %v\n", b, pb)
	// Nota: En los Printf, se pueden usar secuencias de caracteres especiales en el string que serán reemplazados
	// con los valores de las variables. En este caso, el '%v' significa que ahí se escribirá el valor con el
	// tipo de dato de la variable.

	// El símbolo '*' en una variable declarada como puntero, indica que se quiere obtener el valor en vez de
	// la dirección de memoria.
	r := *pa + *pb

	fmt.Printf("La suma del valor de pa y del valor de pb es %v\n", r)
}