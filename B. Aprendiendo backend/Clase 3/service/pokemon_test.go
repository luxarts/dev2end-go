package service

import (
	"clase3/domain"
	"clase3/repository"
	"errors"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestVersus_SquirtleWinsAgainstCharmander(t *testing.T){
	// Given
	charmander := &domain.Pokemon{
		Name: "charmander",
		Type: "fire",
	}
	squirtle := &domain.Pokemon{
		Name: "squirtle",
		Type: "water",
	}

	fire := domain.PokeAPIGetTypeResponse{
		DamageRelations: domain.DamageRelations{
			DoubleDamageFrom: []domain.TypeName{{ "ground"}, { "rock"}, {"water"}},
			DoubleDamageTo:   []domain.TypeName{{ "bug"}, { "steel"}, { "grass"}, { "ice"}},
			HalfDamageFrom:   []domain.TypeName{{ "bug"}, { "steel"}, { "fire"}, { "grass"}, { "ice"}, { "fairy"}},
			HalfDamageTo:     []domain.TypeName{{ "rock"}, { "fire"}, { "water"}, {"dragon"}},
			NoDamageFrom:     []domain.TypeName{},
			NoDamageTo:       []domain.TypeName{},
		},
	}

	pokeAPIRepo := new(repository.PokeAPIRepositoryMock)
	pokeAPIRepo.On("GetType", charmander.Type).Return(&fire, nil)
	pokeSrv := NewPokemonService(pokeAPIRepo)

	// When
	winner, err := pokeSrv.versus(charmander, squirtle)

	// Then
	require.Nil(t, err)
	require.NotNil(t, winner)
	require.Equal(t, squirtle, winner)
	pokeAPIRepo.AssertExpectations(t)
}
func TestVersus_RepoError(t *testing.T){
	// Given
	charmander := &domain.Pokemon{
		Name: "charmander",
		Type: "fire",
	}
	squirtle := &domain.Pokemon{
		Name: "squirtle",
		Type: "water",
	}

	pokeAPIRepo := new(repository.PokeAPIRepositoryMock)
	pokeAPIRepo.On("GetType", charmander.Type).Return(nil, errors.New("error"))
	pokeSrv := NewPokemonService(pokeAPIRepo)

	// When
	winner, err := pokeSrv.versus(charmander, squirtle)

	// Then
	require.Nil(t, winner)
	require.NotNil(t, err)
	require.Equal(t, "error", err.Error())
	pokeAPIRepo.AssertExpectations(t)
}
func TestVersus_Poke1WinsWithDoubleDamage(t *testing.T){
	// Given
	poke1 := &domain.Pokemon{
		Name: "poke1",
		Type: "type1",
	}
	poke2 := &domain.Pokemon{
		Name: "poke2",
		Type: "type2",
	}

	type1 := domain.PokeAPIGetTypeResponse{
		DamageRelations: domain.DamageRelations{
			DoubleDamageFrom: []domain.TypeName{},
			DoubleDamageTo:   []domain.TypeName{{"type2"}},
			HalfDamageFrom:   []domain.TypeName{},
			HalfDamageTo:     []domain.TypeName{},
			NoDamageFrom:     []domain.TypeName{},
			NoDamageTo:       []domain.TypeName{},
		},
	}

	pokeAPIRepo := new(repository.PokeAPIRepositoryMock)
	pokeAPIRepo.On("GetType", poke1.Type).Return(&type1, nil)
	pokeSrv := NewPokemonService(pokeAPIRepo)

	// When
	winner, err := pokeSrv.versus(poke1, poke2)

	// Then
	require.Nil(t, err)
	require.NotNil(t, winner)
	require.Equal(t, poke1, winner)
	pokeAPIRepo.AssertExpectations(t)
}
func TestVersus_NoWinner(t *testing.T){
	// Given
	poke1 := &domain.Pokemon{
		Name: "poke1",
		Type: "type1",
	}
	poke2 := &domain.Pokemon{
		Name: "poke2",
		Type: "type2",
	}

	type1 := domain.PokeAPIGetTypeResponse{
		DamageRelations: domain.DamageRelations{
			DoubleDamageFrom: []domain.TypeName{},
			DoubleDamageTo:   []domain.TypeName{},
			HalfDamageFrom:   []domain.TypeName{},
			HalfDamageTo:     []domain.TypeName{},
			NoDamageFrom:     []domain.TypeName{},
			NoDamageTo:       []domain.TypeName{},
		},
	}

	pokeAPIRepo := new(repository.PokeAPIRepositoryMock)
	pokeAPIRepo.On("GetType", poke1.Type).Return(&type1, nil)
	pokeSrv := NewPokemonService(pokeAPIRepo)

	// When
	winner, err := pokeSrv.versus(poke1, poke2)

	// Then
	require.Nil(t, winner)
	require.Nil(t, err)
	pokeAPIRepo.AssertExpectations(t)
}