package models

type User struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Branch     string `json:"branch"`
	Year       string `json:"year"`
	Password   string `json:"-"`
	IsVerified bool   `json:"is_verified"`
}
