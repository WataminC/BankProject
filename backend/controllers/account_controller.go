package controllers

import (
	"bank/global"
	"bank/models"
	"net/http"
	"strconv"

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
		Money string `json:"money"`
	}

	if err := ctx.ShouldBindBodyWithJSON(&input); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": "Can't read input"})
		return
	}

	money, err := strconv.ParseFloat(input.Money, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": "Can't read input"})
		return
	}

	global.DB.Model(&account).Update("balance", account.Balance+money)

	ctx.JSON(http.StatusOK, gin.H{"New Balance": account.Balance + money})
}
