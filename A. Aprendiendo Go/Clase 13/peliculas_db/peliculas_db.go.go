package peliculas_db

import "errors"

type Pelicula struct {
	ID int
	Titulo string
	Genero string
}

type PeliculasDB interface {
	ObtenerPorTitulo(titulo string) (Pelicula, error)
	ObtenerPorID(id int) (Pelicula, error)
}

type peliculasDB struct {
	tabla map[int]Pelicula
}

func NewPeliculas() PeliculasDB {
	t := make(map[int]Pelicula)

	t[1234] = Pelicula{
		ID: 1234,
		Titulo: "Titanic",
		Genero: "Drama",
	}

	t[456] = Pelicula{
		ID: 456,
		Titulo: "Volver al Futuro",
		Genero: "Ciencia Ficcion",
	}

	t[789] = Pelicula{
		ID: 789,
		Titulo: "Harry Potter",
		Genero: "Fantasia",
	}

	return &peliculasDB{
		tabla: t,
	}
}

func (pdb *peliculasDB) ObtenerPorTitulo(titulo string) (Pelicula, error) {
	// Acá deberíamos conectarnos a una base de datos e ir a buscar la pelicula con una query como
	// SELECT * FROM `peliculas` WHERE `titulo`=titulo;

	for _, p := range pdb.tabla {
		if titulo == p.Titulo {
			return p, nil
		}
	}


	return Pelicula{}, errors.New("no se encontró la película")
}

func (pdb *peliculasDB) ObtenerPorID(id int) (Pelicula, error) {
	// Acá deberíamos conectarnos a una base de datos e ir a buscar la película con una query como
	// SELECT * FROM `peliculas` WHERE `id`=id;
 	p, exists := pdb.tabla[id]

 	if !exists {
 		return p, errors.New("no se encontró la película")
	}

	return pdb.tabla[id], nil
}