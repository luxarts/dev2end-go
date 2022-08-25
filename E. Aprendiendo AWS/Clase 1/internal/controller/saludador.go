package controller

import (
	"clase1/internal/domain"
	"clase1/internal/service"
	"context"
	"net/http"
)

type SaludadorController interface {
	Saludar(ctx context.Context, req domain.Request) (domain.Response, error)
}

type saludadorController struct {
	srv service.SaludadorService
}

func NewSaludadorController(srv service.SaludadorService) SaludadorController {
	return &saludadorController{srv: srv}
}

func (c *saludadorController) Saludar(ctx context.Context, req domain.Request) (domain.Response, error) {
	// Validación de datos de entrada
	if valid, err := req.IsValid(); !valid && err != nil {
		return domain.Response{}, err
	}

	// Ejecuta el servicio
	frase, err := c.srv.Saludar(req.Nombre)
	if err != nil {
		return domain.Response{}, err
	}

	// Arma respuesta
	resp := domain.Response{
		Status: http.StatusOK,
		Data:   frase,
	}

	return resp, nil
}
