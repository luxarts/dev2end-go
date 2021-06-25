package main

import "fmt"

func main() {
	var a int = 10 // Ejemplo: dirección de a = 0x1234
	var b int = 20 // Ejemplo: dirección de b = 0x5678

	// El símbolo '&' indica que se quiere obtener la dirección de la variable en vez de su valor
	// El símbolo '*' indica que lo que se guardará es una dirección
	var pa *int = &a // Ejemplo: pa = 0x1234
	var pb *int = &b // Ejemplo: pb = 0x5678

	fmt.Printf("a vale %v y su dirección es %v\n", a, pa)
	fmt.Printf("b vale %v y su dirección es %v\n", b, pb)

	// El símbolo '*' indica que se quiere obtener un valor en vez de la dirección
	r := *pa + *pb

	fmt.Printf("La suma del valor de pa y del valor de pb es %v\n", r)
}
