package controller

import (
	"clase2/domain"
	"clase2/service"
	"clase2/utils/jsend"
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
		return
	}

	response, err := c.usersService.Create(&userBody)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, jsend.NewResponse(response))
}
func (c *usersController) GetByID(ctx *gin.Context){
	id := ctx.Param("userID")

	if id == "" {
		ctx.JSON(http.StatusBadRequest, "Missing user ID.")
		return
	}

	response, err := c.usersService.GetByID(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if response == nil {
		ctx.JSON(http.StatusNotFound, jsend.NewFailure("not-found"))
		return
	}

	ctx.JSON(http.StatusOK, jsend.NewResponse(response))
}