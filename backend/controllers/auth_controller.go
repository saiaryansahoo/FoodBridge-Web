package controllers

import (
	"backend/config"
	"backend/models"
	"backend/utils"
	"net/http"
	"time"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Secret key for JWT (should be stored in .env)
var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

// Register function
func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.Password = utils.HashPassword(user.Password)
	config.DB.Create(&user)
	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully!"})
}

// Login function
func Login(c *gin.Context) {
	var user models.User
	var input models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Where("email = ?", input.Email).First(&user)
	if !utils.CheckPasswordHash(input.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// ✅ Generate JWT Token
	token, err := GenerateJWT(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	// ✅ Return token in response
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful!",
		"token":   token,
	})
}

// ✅ Function to Generate JWT Token
func GenerateJWT(userID uint) (string, error) {
	// Define token claims
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
