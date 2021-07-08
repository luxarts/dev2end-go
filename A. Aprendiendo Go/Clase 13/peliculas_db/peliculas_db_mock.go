package peliculas_db

import "github.com/stretchr/testify/mock"

type PeliculasDBMock struct{
	mock.Mock
}

func (m *PeliculasDBMock) ObtenerPorTitulo(titulo string) (Pelicula, error) {
	args := m.Called(titulo)

	if args.Get(1) == nil {
		return args.Get(0).(Pelicula), nil
	}

	return args.Get(0).(Pelicula), args.Get(1).(error)
}

func (m *PeliculasDBMock) ObtenerPorID(id int) (Pelicula, error) {
	args := m.Called(id)

	if args.Get(1) == nil {
		return args.Get(0).(Pelicula), nil
	}

	return args.Get(0).(Pelicula), args.Get(1).(error)
}