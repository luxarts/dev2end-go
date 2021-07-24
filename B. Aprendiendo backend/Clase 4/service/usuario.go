package service

import (
	"clase4/domain"
	"clase4/repository"
	"math/rand"
)

type UsuarioService interface {
	Crear(usuario domain.UsuarioPOST) domain.UsuarioDTO
	ObtenerPorID(id uint64) domain.UsuarioDTO
}

type usuarioService struct {
	usuarioRepository repository.UsuarioRepository
}

func NewUsuarioService(usuarioRepository repository.UsuarioRepository) UsuarioService {
	return &usuarioService{
		usuarioRepository: usuarioRepository,
	}
}

func (s *usuarioService) Crear(usuarioPOST domain.UsuarioPOST) domain.UsuarioDTO{
	var u domain.Usuario

	u.FromUsuarioPOST(usuarioPOST)

	u.ID = rand.Uint64()

	s.usuarioRepository.Agregar(u)

	var uDTO domain.UsuarioDTO

	uDTO.FromUsuario(u)

	return uDTO
}

func (s *usuarioService) ObtenerPorID(id uint64) domain.UsuarioDTO {
	u := s.usuarioRepository.ObtenerPorID(id)

	var uDTO domain.UsuarioDTO

	uDTO.FromUsuario(u)

	return uDTO
}