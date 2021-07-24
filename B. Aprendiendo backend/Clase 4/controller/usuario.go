package controller

import (
	"clase4/domain"
	"clase4/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UsuarioController interface {
	Crear(ctx *gin.Context)
	Obtener(ctx *gin.Context)
}

type usuarioController struct {
	userService service.UsuarioService
}

func NewUsuarioController(userService service.UsuarioService) UsuarioController {
	return &usuarioController{
		userService: userService,
	}
}

func (c *usuarioController) Crear(ctx *gin.Context){
	var u domain.UsuarioPOST

	ctx.ShouldBindJSON(&u)

	if !u.EsValido() {
		ctx.String(http.StatusBadRequest, "Body incorrecto")
		return
	}

	uDTO := c.userService.Crear(u)

	ctx.JSON(http.StatusCreated, uDTO)
}
func (c *usuarioController) Obtener(ctx *gin.Context){
	id := ctx.Param("id")

	idInt, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		ctx.String(http.StatusBadRequest, "ID invalido")
		return
	}

	u := c.userService.ObtenerPorID(idInt)

	ctx.JSON(http.StatusOK, u)
}