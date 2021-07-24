package repository

import (
	"clase3/domain"
	"github.com/stretchr/testify/mock"
)

type PokeAPIRepositoryMock struct {
	mock.Mock
}

func (m *PokeAPIRepositoryMock)GetPokemonByName(name string) (*domain.PokeAPIGetPokemonByNameResponse, error){
	args := m.Called(name)

	resp := args.Get(0)
	err := args.Get(1)

	if resp != nil && err == nil{
		return resp.(*domain.PokeAPIGetPokemonByNameResponse), nil
	} else if err != nil && resp == nil{
		return nil, err.(error)
	} else if resp != nil && err != nil {
		return resp.(*domain.PokeAPIGetPokemonByNameResponse), err.(error)
	}

	return nil, nil
}
func (m *PokeAPIRepositoryMock)GetType(name string) (*domain.PokeAPIGetTypeResponse, error) {
	args := m.Called(name)

	resp := args.Get(0)
	err := args.Get(1)

	if resp != nil && err == nil{
		return resp.(*domain.PokeAPIGetTypeResponse), nil
	} else if err != nil && resp == nil{
		return nil, err.(error)
	} else if resp != nil && err != nil {
		return resp.(*domain.PokeAPIGetTypeResponse), err.(error)
	}

	return nil, nil
}