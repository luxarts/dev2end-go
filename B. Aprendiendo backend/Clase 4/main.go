package main

import (
	"clase4/controller"
	"clase4/repository"
	"clase4/service"
	"github.com/gin-gonic/gin"
)

func main(){
	router := InitRouter()
	router.Run("127.0.0.1:9090")
}

func InitRouter() *gin.Engine{
	router := gin.Default()

	MapRoutes(router)

	return router
}

func MapRoutes(r *gin.Engine){
	usrRepo := repository.NewUsuarioRepository()
	usrSrv := service.NewUsuarioService(usrRepo)
	usrCtrl := controller.NewUsuarioController(usrSrv)

	r.GET("/usuario/:id", usrCtrl.Obtener)
	r.POST("/usuario", usrCtrl.Crear)
}



