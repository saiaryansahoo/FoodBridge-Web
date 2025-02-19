package routes

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

func ContactRoutes(r *gin.Engine) {
	contact := r.Group("/contact")
	{
		contact.POST("", controllers.ContactForm) // ✅ Correct API route
	}
}
