package controllers

import (
	"context"

	"projecto-backend/database"
	"projecto-backend/models"

	"github.com/gin-gonic/gin"
)

func GetProfile(c *gin.Context) {
	userId := c.GetInt("user_id")

	var user models.User
	err := database.DB.QueryRow(
		context.Background(),
		"SELECT id, name, email, phone, branch, year FROM users WHERE id=$1",
		userId,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Phone,
		&user.Branch,
		&user.Year,
	)

	if err != nil {
		c.JSON(404, gin.H{"error": "user not found"})
		return
	}

	c.JSON(200, user)
}
