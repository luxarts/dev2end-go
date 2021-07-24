package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := InitRouter()
	router.Run("127.0.0.1:9090")
}

func InitRouter() *gin.Engine {
	router := gin.Default()

	MapRoutes(router)

	return router
}


func MapRoutes(router *gin.Engine) {
	router.GET("/saludar/:nombre", saludar)
}

func saludar(ctx *gin.Context){
	ctx.String(http.StatusOK, "Hola "+ctx.Param("nombre"))
}