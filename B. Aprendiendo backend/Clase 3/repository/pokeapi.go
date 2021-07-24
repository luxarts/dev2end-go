package repository

import (
	"clase3/domain"
	"encoding/json"
	"github.com/go-resty/resty/v2"
)

const (
	pokeAPIBaseURL = "https://pokeapi.co"
	pokeAPIGetPokemonByNameURL = pokeAPIBaseURL+"/api/v2/pokemon/{name}"
	pokeAPIGetTypeURL = pokeAPIBaseURL+"/api/v2/type/{name}"
)

type PokeAPIRepository interface {
	GetPokemonByName(name string) (*domain.PokeAPIGetPokemonByNameResponse, error)
	GetType(name string) (*domain.PokeAPIGetTypeResponse, error)
}

type pokeAPIRepository struct {
	client *resty.Client
}

func NewPokeAPIRepository(rc *resty.Client) PokeAPIRepository{
	return &pokeAPIRepository{
		client: rc,
	}
}

func (r *pokeAPIRepository)GetPokemonByName(name string) (*domain.PokeAPIGetPokemonByNameResponse, error){
	req := r.client.R()
	req = req.SetPathParam("name", name)
	resp, err := req.Get(pokeAPIGetPokemonByNameURL)

	if err != nil {
		return nil, err
	}

	var pokeAPIResponse domain.PokeAPIGetPokemonByNameResponse

	err = json.Unmarshal(resp.Body(), &pokeAPIResponse)

	if err != nil {
		return nil, err
	}

	return &pokeAPIResponse, nil
}
func (r *pokeAPIRepository)GetType(name string) (*domain.PokeAPIGetTypeResponse, error){
	req := r.client.R()
	req = req.SetPathParam("name", name)
	resp, err := req.Get(pokeAPIGetTypeURL)

	if err != nil {
		return nil, err
	}

	var pokeAPIResponse domain.PokeAPIGetTypeResponse

	err = json.Unmarshal(resp.Body(), &pokeAPIResponse)

	if err != nil {
		return nil, err
	}

	return &pokeAPIResponse, nil
}