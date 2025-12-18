package service

import (
	"net/http"
	"payment-sendbox/service/internal/model/response"

	"github.com/gin-gonic/gin"
)

func RegisterUser(context *gin.Context) {
	//todo create registration logic
	success := response.Success{
		Success: true,
		Error:   "",
	}
	context.JSON(http.StatusOK, success)
}

func ChangePassword(context *gin.Context) {
	//todo create password change logic
	success := response.Success{
		Success: true,
		Error:   "",
	}
	context.JSON(http.StatusOK, success)
}
