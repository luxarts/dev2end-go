package controller

import (
	"clase3/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PokemonController interface {
	Fight(ctx *gin.Context)
}

type pokemonController struct {
	pokemonService service.PokemonService
}

func NewPokemonController(pkmnSrv service.PokemonService) PokemonController{
	return &pokemonController{
		pokemonService: pkmnSrv,
	}
}

func (c *pokemonController)Fight(ctx *gin.Context){
	pokemonNames := ctx.QueryArray("name")

	if len(pokemonNames) != 2 {
		ctx.String(http.StatusBadRequest, "Solo pueden pelear dos pokemones a la vez.")
		return
	}

	pokemon, err := c.pokemonService.Fight(pokemonNames)

	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, pokemon.Name + " ganó.")
}