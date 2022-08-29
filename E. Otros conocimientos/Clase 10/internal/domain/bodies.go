package domain

import (
	"errors"
	"strings"
)

const (
	invalidCharacters = "1234567890"
)

type Request struct {
	Nombre string `json:"nombre"`
}

func (req Request) IsValid() (bool, error) {
	if req.Nombre == "" {
		return false, errors.New("empty name")
	}
	if strings.ContainsAny(req.Nombre, invalidCharacters) {
		return false, errors.New("invalid characters in name")
	}

	return true, nil
}

type Response struct {
	Status int    `json:"status"`
	Data   string `json:"data"`
}
