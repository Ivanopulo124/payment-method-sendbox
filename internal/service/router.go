package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct{}

func RegisterRoutes(r *gin.Engine) {

	auth := r.Group("/auth")
	auth.POST("/register", RegisterUser)
	auth.POST("/changePassword", ChangePassword)

	payment := r.Group("/payment")
	payment.POST("/create", Create)
	payment.POST("/status", GetStatus)
	payment.POST("/cancel", Cancel)
	payment.POST("/capture", Cancel)
	payment.POST("/refund", Cancel)

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	})
}
