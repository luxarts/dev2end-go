package service

import (
	"clase1/internal/defines"
	"errors"
	"fmt"
)

type SaludadorService interface {
	Saludar(nombre string) (string, error)
}

type saludadorServiceEs struct {
}
type saludadorServiceEn struct {
}

func NewSaludador(idioma defines.Idioma) (SaludadorService, error) {
	var srv SaludadorService

	// Reemplaza la implementación de acuerdo al idioma
	switch idioma {
	case defines.IdiomaEs:
		srv = &saludadorServiceEs{}
	case defines.IdiomaEn:
		srv = &saludadorServiceEn{}
	default:
		return nil, errors.New("unknown language")
	}

	return srv, nil
}

func (srv *saludadorServiceEs) Saludar(nombre string) (string, error) {
	return fmt.Sprintf("Hola %s!", nombre), nil
}

func (srv *saludadorServiceEn) Saludar(nombre string) (string, error) {
	return fmt.Sprintf("Hello %s!", nombre), nil
}
