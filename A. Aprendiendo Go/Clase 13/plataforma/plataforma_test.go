package plataforma

import (
	"clase13/peliculas_db"
	"errors"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPlataforma_ObtenerPeliculaPorTitulo_OK(t *testing.T) {
	// Given
	peliSpiderman := peliculas_db.Pelicula{
		ID:     999,
		Titulo: "Spiderman",
		Genero: "Accion",
	}

	peliculasDBMock := new(peliculas_db.PeliculasDBMock)
	peliculasDBMock.On("ObtenerPorTitulo", "Spiderman").Return(peliSpiderman, nil)

	plat := NewPlataforma(peliculasDBMock)

	// When
	peli := plat.ObtenerPeliculaPorTitulo("Spiderman")

	// Then
	require.Equal(t, "Spiderman", peli.Titulo)
	require.Equal(t, "Accion", peli.Genero)
	require.Equal(t, 999, peli.ID)
}
func TestPlataforma_ObtenerPeliculaPorTitulo_ErrorSiNoExiste(t *testing.T) {
	// Given
	peliMock := peliculas_db.Pelicula{}

	peliculasDBMock := new(peliculas_db.PeliculasDBMock)
	peliculasDBMock.On("ObtenerPorTitulo", "Spiderman").Return(peliMock, errors.New("no se encontró la película"))

	plat := NewPlataforma(peliculasDBMock)

	// When

	// Then
	require.Panics(t, func(){ plat.ObtenerPeliculaPorTitulo("Spiderman") })
}

func TestPlataforma_ObtenerPeliculaPorID_OK(t *testing.T) {
	// Given
	peliSpiderman := peliculas_db.Pelicula{
		ID:     999,
		Titulo: "Spiderman",
		Genero: "Accion",
	}

	peliculasDBMock := new(peliculas_db.PeliculasDBMock)
	peliculasDBMock.On("ObtenerPorID", 999).Return(peliSpiderman, nil)

	plat := NewPlataforma(peliculasDBMock)

	// When
	msg := plat.ReproducirPeliculaPorID(999)

	// Then
	require.Equal(t, "Reproduciendo Spiderman...", msg)
}
func TestPlataforma_ObtenerPeliculaPorID_ErrorSiNoExiste(t *testing.T) {
	// Given
	peliMock := peliculas_db.Pelicula{}

	peliculasDBMock := new(peliculas_db.PeliculasDBMock)
	peliculasDBMock.On("ObtenerPorID", 999).Return(peliMock, errors.New("no se encontró la película"))

	plat := NewPlataforma(peliculasDBMock)

	// When

	// Then
	require.Panics(t, func(){plat.ReproducirPeliculaPorID(999)})
}