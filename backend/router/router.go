package router

import (
	"bank/controllers"
	"bank/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetRouter() *gin.Engine {
	r := gin.Default()

	// 设置 CORS 中间件
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},                   // 允许访问的前端地址
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},            // 允许的方法
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // 允许的请求头
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true, // 是否允许携带凭证
	}))

	auth := r.Group("/api/auth")
	{
		auth.POST("/login", controllers.Login)
		auth.POST("/register", controllers.Register)
	}

	api := r.Group("/api")
	api.Use(middlewares.AuthMiddleWare())
	{
		api.GET("/account", controllers.GetInfo)
		api.POST("/transaction/deposit", controllers.Deposit)
		api.POST("/transaction/withdraw", controllers.Withdraw)
		api.POST("/transaction/transfer", controllers.AddTransaction)
		api.GET("/transaction", controllers.GetTransaction)
	}

	return r
}
