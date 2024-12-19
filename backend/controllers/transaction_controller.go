package controllers

import (
	"bank/global"
	"bank/models"

	// "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
 * Add Transaction API
 *
 * Input:
 * {
 *   "account_id": "string",
 *   "target_account_number": "string",
 *   "amount": 0
 * }
 *
 * Output:
 * {
 *   "message": "string"
 * }
 *
 */

func AddTransaction(ctx *gin.Context) {
	_, exists := ctx.Get("username")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"Error": "Unauthorized access"})
		return
	}

	var input struct {
		AccountID           uint    `json:"account_id"`
		TargetAccountNumber string  `json:"target_account_number"`
		Amount              float64 `json:"amount"`
	}
	// var input models.Transaction

	// fmt.Printf("input: %v\n", input)

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	// Check if account exists
	if err := global.DB.Where("id = ?", input.AccountID).First(&models.Account{}).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Account not found"})
		return
	}

	// Check if target account exists
	if err := global.DB.Where("account_number = ?", input.TargetAccountNumber).First(&models.Account{}).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Target account not found"})
		return
	}

	// Check if account has enough balance
	if err := global.DB.Where("id = ? AND balance >= ?", input.AccountID, input.Amount).First(&models.Account{}).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Insufficient balance"})
		return
	}

	// Start transaction
	tx := global.DB.Begin()

	// Recover from panic
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var transaction models.Transaction
	transaction.AccountID = input.AccountID
	transaction.TargetAccountNumber = input.TargetAccountNumber
	transaction.Amount = input.Amount
	transaction.Type = "Transfer"

	// Add transaction
	if err := tx.Create(&transaction).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		tx.Rollback()
		return
	}

	// Update source account balance
	var sourceAccount models.Account
	if err := tx.Where("id = ?", input.AccountID).First(&sourceAccount).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err})
		tx.Rollback()
		return
	}
	sourceAccount.Balance -= input.Amount

	if err := tx.Save(&sourceAccount).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		tx.Rollback()
		return
	}

	// Update target account balance
	var targetAccount models.Account
	if err := tx.Where("account_number = ?", input.TargetAccountNumber).First(&targetAccount).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err})
		tx.Rollback()
		return
	}
	targetAccount.Balance += input.Amount

	if err := tx.Save(&targetAccount).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error})
		tx.Rollback()
		return
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	// Return success message
	ctx.JSON(http.StatusOK, gin.H{"message": "Transaction successful"})
}
