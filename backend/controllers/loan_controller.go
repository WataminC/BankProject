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

	c.JSON(http.StatusOK, response)
}

func ProcessLoanRequest(c *gin.Context) {
	var input struct {
		ID       uint `json:"id"`
		Approved bool `json:"is_approved"`
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

	var status = "approved"
	if !input.Approved {
		status = "rejected"
	}

	loanRequest.Status = status
	if err := tx.Save(&loanRequest).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		tx.Rollback()
		return
	}

	if !input.Approved {
		tx.Commit()
		c.JSON(http.StatusOK, gin.H{"message": "Loan request rejected"})
		return
	}

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
	if err := global.DB.Where("user_id = ? AND (status = ? OR status = ?)", input.UserID, "approved", "rejected").First(&loanRequest).Error; err != nil {
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

	if loanRequest.Status == "rejected" {
		ctx.JSON(http.StatusOK, gin.H{"is_succeed": false})
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

func RepayLoan(ctx *gin.Context) {
	var input struct {
		UserID uint    `json:"user_id"`
		Amount float64 `json:"amount"`
	}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userAccount models.Account
	if err := global.DB.Where("user_id = ?", input.UserID).First(&userAccount).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"is_succeed": false, "message": "account not found"})
		return
	}

	var userLoan models.Loan
	if err := global.DB.Where("user_id = ?", input.UserID).First(&userLoan).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"is_succeed": false, "message": "loan not found"})
		return
	}

	if userAccount.Balance < input.Amount {
		ctx.JSON(http.StatusOK, gin.H{"is_succeed": false, "message": "insufficient balance"})
		return
	}

	tx := global.DB.Begin()

	if input.Amount >= userLoan.TotalAmount {
		repayAll := userLoan.TotalAmount
		userAccount.Balance -= repayAll
		if err := tx.Save(&userAccount).Error; err != nil {
			tx.Rollback()
			ctx.JSON(http.StatusInternalServerError, gin.H{"is_succeed": false, "message": err.Error()})
			return
		}
		if err := tx.Unscoped().Delete(&userLoan).Error; err != nil {
			tx.Rollback()
			ctx.JSON(http.StatusInternalServerError, gin.H{"is_succeed": false, "message": err.Error()})
			return
		}
		tx.Commit()
		ctx.JSON(http.StatusOK, gin.H{"is_succeed": true, "message": "repay successfully"})
		return
	}

	userAccount.Balance -= input.Amount
	userLoan.TotalAmount -= input.Amount

	if err := tx.Save(&userAccount).Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, gin.H{"is_succeed": false, "message": err.Error()})
		return
	}

	if err := tx.Save(&userLoan).Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, gin.H{"is_succeed": false, "message": err.Error()})
		return
	}

	tx.Commit()
	ctx.JSON(http.StatusOK, gin.H{"is_succeed": true, "message": "repay successfully"})
}