package plataforma

import (
	"clase13/peliculas_db"
)

type Plataforma interface {
	ObtenerPeliculaPorTitulo(titulo string) peliculas_db.Pelicula
	ReproducirPeliculaPorID(id int) string
}

type plataforma struct {
	peliculasDB peliculas_db.PeliculasDB
}

func NewPlataforma(peliculasDB peliculas_db.PeliculasDB) Plataforma {
	return &plataforma{
		peliculasDB: peliculasDB,
	}
}

func (p *plataforma) ObtenerPeliculaPorTitulo(titulo string) peliculas_db.Pelicula {
	peli, err := p.peliculasDB.ObtenerPorTitulo(titulo)

	if err != nil {
		panic(err)
	}

	return peli
}
func (p *plataforma) ReproducirPeliculaPorID(id int) string {
	peli, err := p.peliculasDB.ObtenerPorID(id)

	if err != nil {
		panic(err)
	}

	return "Reproduciendo "+peli.Titulo+"..."
}