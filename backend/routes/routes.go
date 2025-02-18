package routes

import "github.com/gin-gonic/gin"

func SetupRoutes(r *gin.Engine) {
	AuthRoutes(r)
	FoodRoutes(r)
	ContactRoutes(r)
}
