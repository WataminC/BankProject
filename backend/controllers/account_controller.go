package controllers

import (
	"bank/global"
	"bank/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetInfo(ctx *gin.Context) {
	var userAccountInfo struct {
		UserID        uint    `gorm:"column:user_id"`
		AccountNumber string  `gorm:"column:account_number"`
		Balance       float64 `gorm:"column:balance"`
		Name          string  `gorm:"column:name"`
		Email         string  `gorm:"column:email"`
	}

	username, exists := ctx.Get("username")

	if !exists {
		// 处理未找到的情况
		ctx.JSON(http.StatusUnauthorized, gin.H{"Error": "Unauthorized access"})
		return
	}

	global.DB.Raw(`
		SELECT *
		FROM accounts JOIN users ON accounts.user_id = users.id
		WHERE users.name = ?
	`, username).Scan(&userAccountInfo)

	ctx.JSON(http.StatusOK, gin.H{
		"userID":         userAccountInfo.UserID,
		"username":       userAccountInfo.Name,
		"email":          userAccountInfo.Email,
		"account_number": userAccountInfo.AccountNumber,
		"balance":        userAccountInfo.Balance})
}

func Deposit(ctx *gin.Context) {
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

	var input struct {
		Amount string `json:"amount"`
	}

	if err := ctx.ShouldBindBodyWithJSON(&input); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": "Invalid input"})
		return
	}

	money, err := strconv.ParseFloat(input.Amount, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": "Invalid input"})
		return
	}

	global.DB.Model(&account).Update("balance", account.Balance+money)

	ctx.JSON(http.StatusOK, gin.H{
		"message":     "Deposit successfully",
		"new_balance": account.Balance})
}

func Withdraw(ctx *gin.Context) {
	// 从上下文中获取用户名
	username, exists := ctx.Get("username")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
		return
	}

	var account models.Account

	global.DB.Raw(`
		SELECT accounts.*
		FROM accounts JOIN users ON accounts.user_id = users.id
		WHERE users.name = ?
	`, username).Scan(&account)

	var input struct {
		Amount string `json:"amount"`
	}

	if err := ctx.ShouldBindBodyWithJSON(&input); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": "Invalid input"})
		return
	}

	money, err := strconv.ParseFloat(input.Amount, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": "Invalid input"})
		return
	}

	// 检查账户余额是否足够
	if account.Balance < money {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient balance"})
		return
	}

	global.DB.Model(&account).Update("balance", account.Balance-money)

	ctx.JSON(http.StatusOK, gin.H{
		"message":     "Withdraw successfully",
		"new_balance": account.Balance})
}
