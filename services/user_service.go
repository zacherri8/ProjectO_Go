package services

import (
	"context"
	"projecto-backend/database"
	"projecto-backend/models"
)

func GetUserByID(id int) (models.User, error) {
	var user models.User

	err := database.DB.QueryRow(
		context.Background(),
		"SELECT id, name, email, year, branch, is_verified FROM users WHERE id=$1",
		id,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Year,
		&user.Branch,
		&user.IsVerified,
	)

	return user, err
}
