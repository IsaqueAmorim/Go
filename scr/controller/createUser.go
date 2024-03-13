package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/teste/scr/configuration/validation"
	"github.com/teste/scr/controller/model/request"
)

func CreateUser(c *gin.Context) {
	var request request.UserRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		restError := validation.ValidateUserError(err)

		c.JSON(restError.Code, restError)
		return
	}

	fmt.Println(request)
}
