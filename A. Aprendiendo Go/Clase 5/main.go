package main

func main() {
	contador := 0
	condicion := false

	// While
	for condicion == true {
		// Hago algo

		contador++
		if contador == 10 {
			condicion = true
		}
	}

	// For
	for i := 0; i < 10; i++ {
		// Hago algo
	}
}
