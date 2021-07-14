package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := InitRouter()
	router.Run(":8080")
}

func InitRouter() *gin.Engine {
	router := gin.Default()

	MapRoutes(router)

	return router
}


func MapRoutes(router *gin.Engine) {

}