package controllers

import (
	"backend/config"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListFood(c *gin.Context) {
	var food models.Food
	if err := c.ShouldBindJSON(&food); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&food)
	c.JSON(http.StatusCreated, gin.H{"message": "Food listed successfully!"})
}

func GetFood(c *gin.Context) {
	var food []models.Food
	config.DB.Find(&food)
	c.JSON(http.StatusOK, food)
}
