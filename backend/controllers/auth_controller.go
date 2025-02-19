package controllers

import (
	"backend/config"
	"backend/models"
	"backend/utils"
	"errors"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// ✅ Get JWT Secret from Environment
var jwtSecret = []byte(getJWTSecret())

func getJWTSecret() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "default_secret_key" // ⚠️ Change this in production!
	}
	return secret
}

// ✅ Function to Generate JWT Token
func GenerateJWT(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// ✅ Safe Write Function to Handle DB Locks
func SafeDBWrite(action func() error) error {
	for i := 0; i < 5; i++ { // Retry 5 times
		err := action()
		if err != nil && strings.Contains(err.Error(), "database is locked") {
			time.Sleep(time.Millisecond * 100) // Wait and retry
			continue
		}
		return err
	}
	return errors.New("database is locked, retries exhausted")
}

// ✅ Register Function (Fixed Database Lock Issues)
func Register(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if email already exists
	var existingUser models.User
	if err := config.DB.Where("email = ?", input.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Email already registered"})
		return
	}

	// Hash password before saving
	input.Password = utils.HashPassword(input.Password)

	// ✅ Use SafeDBWrite to prevent locking errors
	err := SafeDBWrite(func() error {
		return config.DB.Create(&input).Error
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	// ✅ Generate JWT Token after registration
	token, err := GenerateJWT(input.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	// ✅ Return token in response
	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully!",
		"token":   token,
	})
}

// ✅ Login Function
func Login(c *gin.Context) {
	var input models.User
	var user models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find user by email
	result := config.DB.Where("email = ?", input.Email).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Check password
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

// ✅ Middleware to Protect Routes
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
			c.Abort()
			return
		}

		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set("user_id", claims["user_id"])
		c.Next()
	}
}
