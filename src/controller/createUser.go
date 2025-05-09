package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/nascimentoFabiano/primeira-api-golang/src/configuration/rest_err"
	"github.com/nascimentoFabiano/primeira-api-golang/src/controller/model/request"
)

func CreateUser(c *gin.Context) {

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_err.NewBadRequestError(
			fmt.Sprintf("There are some incorret fields, error=%s\n", err.Error()))

		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userRequest)
}
