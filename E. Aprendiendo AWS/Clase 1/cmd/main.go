package main

import (
	"clase1/internal/controller"
	"clase1/internal/defines"
	"clase1/internal/service"
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
