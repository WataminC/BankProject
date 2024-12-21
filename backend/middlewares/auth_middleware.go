package middlewares

import (
	"bank/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"Error": "Haven't login"})
			ctx.Abort()
			return
		}

		username, err := utils.ParseJWT(token)

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"Error": "Unauthorized user"})
			ctx.Abort()
			return
		}

		ctx.Set("username", username)

		ctx.Next()
	}
}

func AdminMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"Error": "Haven't login"})
			ctx.Abort()
			return
		}

		username, err := utils.ParseJWT(token)

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"Error": "Unauthorized user"})
			ctx.Abort()
			return
		}

		if username != "admin" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"Error": "Unauthorized user"})
			ctx.Abort()
			return
		}

		// ctx.Set("username", username)
		ctx.Next()
	}
}
