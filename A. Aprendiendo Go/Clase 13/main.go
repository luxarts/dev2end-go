package main

import (
	"clase13/peliculas_db"
	"clase13/plataforma"
	"fmt"
)

func main() {
	peliculasDB := peliculas_db.NewPeliculas()

 	netflix := plataforma.NewPlataforma(peliculasDB)

 	p := netflix.ObtenerPeliculaPorTitulo("Harry Potter")

 	fmt.Println(netflix.ReproducirPeliculaPorID(p.ID))
}

