package services

import (
	"projecto-backend/utils"
)

func HashPassword(password string) (string, error) {
	return utils.HashPassword(password)
}

func CheckPassword(hash string, password string) bool {
	return utils.CheckPassword(hash, password)
}
