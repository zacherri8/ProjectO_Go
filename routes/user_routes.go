package routes

import (
	"projecto-backend/controllers"
	"projecto-backend/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	u := r.Group("/user")
	u.Use(middleware.AuthMiddleware())

	u.GET("/profile", controllers.GetProfile)
}
