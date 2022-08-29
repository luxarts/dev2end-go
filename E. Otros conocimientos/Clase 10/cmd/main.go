package main

import (
	"clase10/internal/controller"
	"clase10/internal/defines"
	"clase10/internal/service"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
)

func main() {
	// Instancia servicios
	saludadorSrv, err := service.NewSaludador(defines.IdiomaEs)
	if err != nil {
		log.Fatalln(err)
	}

	// Instancia controladores
	saludadorCtrl := controller.NewSaludadorController(saludadorSrv)

	// Vincula el handler con Lambda
	lambda.Start(saludadorCtrl.Saludar)
}
