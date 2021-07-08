package peliculas_db

import "github.com/stretchr/testify/mock"

// Creamos una estructura que herede las propiedades de mock.Mock
type PeliculasDBMock struct{
	mock.Mock
}

// Creamos los métodos para la estructura para implementar la interface `PeliculasDB`
func (m *PeliculasDBMock) ObtenerPorTitulo(titulo string) (Pelicula, error) {
	// La función `Called` recibe todos los parámetros que reciba nuestra función y devuelve los argumentos que
	// usaremos en los valores de retorno
	args := m.Called(titulo)

	// El método `Get` recibe el índice del valor de retorno de la función. Por ejemplo, si la función devuelve 3
	// valores, Get(0) será el primero, Get(1) el segundo y Get(2) el tercero.
	if args.Get(1) == nil {
		// El método Get(0) devuelve un valor del tipo interface{}, por lo que necesitamos determinar el tipo de dato.
		// Para esto usamos la expresión .(tipo_de_datos) después de la variable.
		// Nota: No se puede determinar un tipo de datos si la variable es nil, por eso necesitamos este if.
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