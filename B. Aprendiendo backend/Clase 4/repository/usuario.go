package repository

import (
	"clase4/domain"
	"encoding/csv"
	"os"
	"strconv"
)

type UsuarioRepository interface {
	Agregar(usuario domain.Usuario)
	ObtenerPorID(id uint64) domain.Usuario
}

type usuarioRepository struct {

}

func NewUsuarioRepository() UsuarioRepository {
	return &usuarioRepository{}
}

func (r *usuarioRepository) Agregar(usuario domain.Usuario) {
	// Abrimos el archivo CSV
	file, err := os.OpenFile("usersdb.csv", os.O_RDWR | os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Escribimos la linea en el CSV
	csvWriter := csv.NewWriter(file)

	row := usuario.ToRow()

	csvWriter.Write(row)

	csvWriter.Flush()
}
func (r *usuarioRepository) ObtenerPorID(id uint64) domain.Usuario {
	idStr := strconv.FormatUint(id, 10)

	// Abrimos el archivo CSV
	file, err := os.Open("usersdb.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Leemos todas las filas
	csvReader := csv.NewReader(file)

	rows, err := csvReader.ReadAll()
	if err != nil {
		panic(err)
	}

	var u domain.Usuario

	for _, row := range rows {
		rowID := row[0]

		if idStr == rowID {
			u.ID = id
			u.Nombre = row[1]
			u.Email = row[2]
			u.Contrasenia = row[3]
			u.FechaDeNacimiento = row[4]
			break
		}
	}

	return u
}