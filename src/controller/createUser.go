package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kasslima/go-crud/src/configuration/rest_err"
	"github.com/kasslima/go-crud/src/controller/model/request"
)

func CreateUser(c *gin.Context) {
	var userResquest request.UserRequest
	if err := c.ShouldBindJSON(&userResquest); err != nil {
		restErr := rest_err.NewBadRequestError(
			fmt.Sprintf("There are some incorrect fields in the request body"))
		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userResquest)
}
