package service

import (
	"clase3/domain"
	"clase3/repository"
)

type PokemonService interface {
	Fight(names []string) (*domain.Pokemon, error)
	versus(poke1 *domain.Pokemon, poke2 *domain.Pokemon) (*domain.Pokemon, error)
}

type pokemonService struct {
	pokeAPIRepository repository.PokeAPIRepository
}

func NewPokemonService(pokeapirepo repository.PokeAPIRepository) PokemonService{
	return &pokemonService{
		pokeAPIRepository: pokeapirepo,
	}
}

func (s *pokemonService)Fight(names []string) (*domain.Pokemon, error){
	var pokemons []domain.Pokemon

	for _, name := range names {
		response, err := s.pokeAPIRepository.GetPokemonByName(name)

		if err != nil {
			return nil, err
		}

		pokemons = append(pokemons, domain.Pokemon{
			Name: name,
			Type: response.Types[0].Type.Name,
		})
	}

	// Versus
	winner, err := s.versus(&pokemons[0], &pokemons[1])

	return winner, err
}

func (s *pokemonService) versus(poke1 *domain.Pokemon, poke2 *domain.Pokemon) (*domain.Pokemon, error){
	damagesForType1, err := s.pokeAPIRepository.GetType(poke1.Type)
	if err != nil {
		return nil, err
	}

	poke2Life := 4

	// Calculamos el daño del poke 1 al poke 2
	if containsType(damagesForType1.DamageRelations.DoubleDamageTo, poke2.Type) {
		// Inflige doble daño al poke 2
		poke2Life-=4
	} else if containsType(damagesForType1.DamageRelations.HalfDamageTo, poke2.Type) {
		// Inflige la mitad de daño al poke 2
		poke2Life-=1
	} else if !containsType(damagesForType1.DamageRelations.NoDamageTo, poke2.Type) {
		// Inflige daño normal
		poke2Life-=2
	}

	poke1Life := 4
	// Calcula el daño del poke 2 al poke 1
	if containsType(damagesForType1.DamageRelations.DoubleDamageFrom, poke2.Type){
		// Recibe doble daño del poke 2
		poke1Life-=4
	} else if containsType(damagesForType1.DamageRelations.HalfDamageFrom, poke2.Type){
		// Recibe la mitad del daño del poke 2
		poke1Life-=1
	} else if !containsType(damagesForType1.DamageRelations.NoDamageFrom, poke2.Type){
		// Recibe daño normal
		poke1Life-=2
	}

	if poke1Life > poke2Life {
		// Si poke 1 queda con mas vida
		return poke1, nil
	} else if poke2Life > poke1Life {
		// Si poke 2 queda con mas vida
		return poke2, nil
	} else {
		// Si quedan con la misma cantidad de vida
		return nil, nil
	}
}

func containsType(types []domain.TypeName, typeName string) bool {
	for _, t := range types {
		if t.Name == typeName {
			return true
		}
	}
	return false
}