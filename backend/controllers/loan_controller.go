package controllers

import (
	"bank/global"
	"bank/models"
	"net/http"

	"github.com/gin-gonic/gin"
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
