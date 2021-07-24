package main

import (
	"clase3/controller"
	"clase3/repository"
	"clase3/service"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

func main(){
	router := NewRouter()
	router.Run("127.0.0.1:9090")
}

func NewRouter() *gin.Engine{
	router := gin.Default()

	mapRoutes(router)

	return router
}

func mapRoutes(r *gin.Engine){
	restClient := resty.New()

	pokeAPIRepo := repository.NewPokeAPIRepository(restClient)
	pokemonService := service.NewPokemonService(pokeAPIRepo)
	pokemonController := controller.NewPokemonController(pokemonService)

	r.GET("/fight", pokemonController.Fight)
}