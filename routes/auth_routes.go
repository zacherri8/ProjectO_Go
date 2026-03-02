package routes

import (
	"projecto-backend/controllers"
	"projecto-backend/middleware"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)

	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware())
	api.GET("/me", controllers.GetProfile)
}
