package controllers

import (
	"bank/global"
	"bank/models"
	"bank/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 用户注册
func Register(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindBodyWithJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// 验证邮箱格式
	if !utils.IsValidEmail(user.Email) {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid input"})
		return
	}

	// 验证邮箱是否存在
	var existingUser models.User
	if err := global.DB.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Email already registered"})
		return
	}

	// 哈希存储密码
	hash, err := utils.HashPassword(user.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": "Can't hash password"})
		return
	}

	user.Password = hash

	// 将用户插入表中
	if err := global.DB.Create(&user).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to create user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Message": "Insert user successfully"})
}

func Login(ctx *gin.Context) {
	var input struct {
		Name string `json:"name"`
		Password string `json:"password"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	var user models.User

	if err := global.DB.Where("name = ?", input.Name).First(&user).Error; err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"Error": "wrong credentials"})
		return
	}

	if !utils.CheckPassword(input.Password, user.Password) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"Error": "wrong credentials"})
		return
	}

	token, err := utils.GenerateJWT(user.Name)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
