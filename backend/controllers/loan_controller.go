package controllers

import (
	"bank/global"
	"bank/models"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func LoanRequest(c *gin.Context) {
	var input struct {
		UserId   uint    `json:"user_id"`
		Amount   float64 `json:"amount"`
		Interest float64 `json:"interest"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	LoanRequest := models.LoanRequest{
		UserID:   input.UserId,
		Amount:   input.Amount,
		Interest: input.Interest,
		Status:   "pending",
	}

	if err := global.DB.Create(&LoanRequest).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message": "Loan request created successfully"})
}

func GetLoanRequest(c *gin.Context) {
	var loanRequest []models.LoanRequest

	if err := global.DB.Find(&loanRequest).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	type responseBody struct {
		ID       uint    `json:"id"`
		UserID   uint    `json:"user_id"`
		Amount   float64 `json:"amount"`
		Interest float64 `json:"interest"`
		Status   string  `json:"status"`
	}

	var response []responseBody
	for _, lr := range loanRequest {
		if lr.Status == "pending" {
			response = append(response, responseBody{
				ID:       lr.ID,
				UserID:   lr.UserID,
				Amount:   lr.Amount,
				Interest: lr.Interest,
				Status:   lr.Status,
			})
		}
	}

	c.JSON(http.StatusOK, gin.H{"loan_request": response})
}

func ApproveLoanRequest(c *gin.Context) {
	var input struct {
		ID uint `json:"id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var loanRequest models.LoanRequest
	if err := global.DB.First(&loanRequest, input.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tx := global.DB.Begin()

	loan := models.Loan{
		UserID:      loanRequest.UserID,
		Amount:      loanRequest.Amount,
		Interest:    loanRequest.Interest,
		TotalAmount: loanRequest.Amount * (1 + 0.01*loanRequest.Interest),
		Status:      "approved",
	}

	if err := tx.Create(&loan).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		tx.Rollback()
		return
	}

	loanRequest.Status = "approved"
	if err := tx.Save(&loanRequest).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		tx.Rollback()
		return
	}

	var userAccount models.Account
	if err := tx.Where("user_id = ?", loan.UserID).First(&userAccount).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		tx.Rollback()
		return
	}

	if err := tx.Model(&userAccount).Update("balance", userAccount.Balance+loan.Amount).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		tx.Rollback()
		return
	}

	tx.Commit()

	c.JSON(http.StatusOK, gin.H{"message": "Loan request approved"})
}

func QueryLoanRequest(ctx *gin.Context) {
	var input struct {
		UserID uint `json:"user_id"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var loanRequest models.LoanRequest
	if err := global.DB.Where("user_id = ? AND status = ?", input.UserID, "approved").First(&loanRequest).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "record not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Delete loan request
	if err := global.DB.Unscoped().Delete(&loanRequest).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"is_succeed": true})
}

func QueryLoan(ctx *gin.Context) {
	var input struct {
		UserID uint `json:"user_id"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var loan models.Loan
	if err := global.DB.Where("user_id = ?", input.UserID).First(&loan).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "record not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var output struct {
		ID          uint    `json:"id"`
		Amount      float64 `json:"amount"`
		Interest    float64 `json:"interest"`
		TotalAmount float64 `json:"total_amount"`
		Status      string  `json:"status"`
	}

	output.ID = loan.ID
	output.Amount = loan.Amount
	output.Interest = loan.Interest
	output.TotalAmount = loan.TotalAmount
	output.Status = loan.Status

	ctx.JSON(http.StatusOK, output)
}