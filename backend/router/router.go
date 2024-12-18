package router

import (
	"bank/controllers"
	"bank/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	auth := r.Group("/api/auth")
	{
		auth.POST("/login", controllers.Login)
		auth.POST("/register", controllers.Register)
	}

	api := r.Group("/api")
	api.Use(middlewares.AuthMiddleWare())
	{
		api.GET("/account", controllers.GetBalance)
		api.POST("/transaction/deposit", controllers.Deposit)
		api.POST("/transaction/withdraw", controllers.Withdraw)
	}

	return r
}
