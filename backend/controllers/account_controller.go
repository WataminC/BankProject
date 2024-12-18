package controllers

import (
	"bank/global"
	"bank/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBalance(ctx *gin.Context) {
	username, exists := ctx.Get("username")

	if !exists {
		// 处理未找到的情况
		ctx.JSON(http.StatusUnauthorized, gin.H{"Error": "Unauthorized access"})
		return
	}

	var account models.Account
	
	global.DB.Raw(`
		SELECT accounts.*
		FROM accounts JOIN users ON accounts.user_id = users.id
		WHERE users.name = ?
	`, username).Scan(&account)

	ctx.JSON(http.StatusOK, gin.H{"balance": account.Balance})
}

func Deposit(ctx *gin.Context) {
	// 从上下文中获取用户名
	username, exists := ctx.Get("username")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
		return
	}

	// 根据用户名查找用户
	var user models.User
	if err := global.DB.Where("name = ?", username).First(&user).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	// 查找用户的账户
	var account models.Account
	if err := global.DB.Where("user_id = ?", user.ID).First(&account).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Account not found"})
		return
	}

	// 解析请求中的存款金额
	var input struct {
		Amount float64 `json:"amount"`
	}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// 检查存款金额的有效性
	if input.Amount <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Deposit amount must be positive"})
		return
	}

	// 更新账户余额
	account.Balance += input.Amount
	if err := global.DB.Save(&account).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update balance"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Deposit successful", "new_balance": account.Balance})
}
