package domain

import (
	"strconv"
	"strings"
)

type UsuarioPOST struct {
	Nombre string `json:"nombre"`
	Email string `json:"email"`
	Contrasenia string `json:"contrasenia"`
	FechaDeNacimiento string `json:"fecha_de_nacimiento"`
}
func (u *UsuarioPOST) EsValido() bool {
	return u.Nombre != "" &&
		u.Contrasenia != "" &&
		u.Email != "" &&
		strings.Contains(u.Email, "@") &&
		u.FechaDeNacimiento != ""
}

type UsuarioDTO struct {
	ID uint64 `json:"id"`
	Nombre string `json:"nombre"`
	Email string `json:"email"`
	FechaDeNacimiento string `json:"fecha_de_nacimiento"`
	Edad uint8 `json:"edad"`
}

func (u *UsuarioDTO) FromUsuario(usuario Usuario) {
	u.ID = usuario.ID
	u.Nombre = usuario.Nombre
	u.Email = usuario.Email
	u.FechaDeNacimiento = usuario.FechaDeNacimiento
}

type Usuario struct {
	ID uint64
	Nombre string
	Email string
	Contrasenia string
	FechaDeNacimiento string
}

func (u *Usuario) FromUsuarioPOST(uPOST UsuarioPOST){
	u.Nombre = uPOST.Nombre
	u.Email = uPOST.Email
	u.Contrasenia = uPOST.Contrasenia
	u.FechaDeNacimiento = uPOST.FechaDeNacimiento
}

func (u *Usuario) ToRow() []string{
	var row []string

	idStr := strconv.FormatUint(u.ID, 10)

	return append(row, idStr, u.Nombre, u.Email, u.Contrasenia, u.FechaDeNacimiento)
}

