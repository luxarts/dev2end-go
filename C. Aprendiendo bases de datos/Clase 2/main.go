package main

import (
	"clase2/controller"
	"clase2/repository"
	"clase2/service"
	"github.com/gin-gonic/gin"
)

func main() {
	router := initRouter()

	router.Run("localhost:9090")
}

func initRouter() *gin.Engine{
	router := gin.Default()

	mapRoutes(router)

	return router
}

func mapRoutes(r *gin.Engine) {
	// Instanciamos el repository
	usersRepo := repository.NewUsersRepository()
	// Instanciamos el service
	usersSrv := service.NewUsersService(usersRepo)
	// Instanciamos el controller
	usersCtrl := controller.NewUsersController(usersSrv)

	r.POST("/user", usersCtrl.Create)
	r.GET("/user/:userID", usersCtrl.GetByID)
	r.DELETE("/user/:userID", usersCtrl.DeleteByID)
}