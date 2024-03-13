package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/teste/scr/configuration/rest_error"
)

func FindUserByID(c *gin.Context) {
	teste := rest_error.NewInternalServerError("ERRO INTERNO, CAIA FORA")
	c.JSON(teste.Code,teste)
}

func FindUserByEmail(c *gin.Context) {}
