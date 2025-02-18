package routes

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

func FoodRoutes(r *gin.Engine) {
	food := r.Group("/food")
	{
		food.POST("/list", controllers.ListFood)
		food.GET("/", controllers.GetFood)
	}
}
