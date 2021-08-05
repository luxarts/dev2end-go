package controller

import (
	"clase2/domain"
	"clase2/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UsersController interface {
	Create(ctx *gin.Context)
	GetByID(ctx *gin.Context)
}

type usersController struct {
	usersService service.UsersService
}

func NewUsersController(srv service.UsersService) UsersController {
	return &usersController{
		usersService: srv,
	}
}

func (c *usersController) Create(ctx *gin.Context){
	var userBody domain.UsersPOSTBody

	if err := ctx.ShouldBindJSON(&userBody); err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid body.")
	}

	response := c.usersService.Create(&userBody)

	ctx.JSON(http.StatusCreated, response)
}
func (c *usersController) GetByID(ctx *gin.Context){
	id := ctx.Param("userID")

	if id == "" {
		ctx.JSON(http.StatusBadRequest, "Missing user ID.")
	}

	response := c.usersService.GetByID(id)

	ctx.JSON(http.StatusOK, response)
}