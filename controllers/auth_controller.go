package controllers

import (
	"context"
	"net/http"
	"strings"

	"projecto-backend/database"
	"projecto-backend/models"
	"projecto-backend/services"
	"projecto-backend/utils"

	"github.com/gin-gonic/gin"
)

type SignupBody struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Branch   string `json:"branch"`
	Year     string `json:"year"`
	Password string `json:"password"`
}

func Signup(c *gin.Context) {
	var body SignupBody
	c.Bind(&body)

	if !strings.HasSuffix(body.Email, "@kiit.ac.in") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "use KIIT email only"})
		return
	}

	hashed, _ := services.HashPassword(body.Password)

	_, err := database.DB.Exec(
		context.Background(),
		`INSERT INTO users (name, email, phone, branch, year, password, is_verified)
		 VALUES ($1,$2,$3,$4,$5,$6,$7)`,
		body.Name, body.Email, body.Phone, body.Branch, body.Year, hashed, false,
	)

	if err != nil {
		c.JSON(400, gin.H{"error": "email already registered"})
		return
	}

	// Create OTP
	otp, _ := services.CreateOTP(body.Email)
	services.SendEmail(body.Email, "Your OTP", "Your OTP is: "+otp)

	c.JSON(200, gin.H{"message": "Signup successful, verify OTP"})
}

type LoginBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var body LoginBody
	c.Bind(&body)

	var user models.User
	err := database.DB.QueryRow(
		context.Background(),
		"SELECT id, name, email, password, is_verified FROM users WHERE email=$1",
		body.Email,
	).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.IsVerified)

	if err != nil {
		c.JSON(400, gin.H{"error": "invalid credentials"})
		return
	}

	if !services.CheckPassword(user.Password, body.Password) {
		c.JSON(400, gin.H{"error": "invalid credentials"})
		return
	}

	if !user.IsVerified {
		c.JSON(403, gin.H{"error": "email not verified"})
		return
	}

	token, _ := utils.GenerateJWT(user.ID)

	c.JSON(200, gin.H{
		"message": "login successful",
		"token":   token,
	})
}

func GetProfile(c *gin.Context) {
	userId := c.GetInt("user_id")

	var user models.User
	err := database.DB.QueryRow(
		context.Background(),
		"SELECT id, name, email, phone, branch, year FROM users WHERE id=$1",
		userId,
	).Scan(&user.ID, &user.Name, &user.Email, &user.Phone, &user.Branch, &user.Year)

	if err != nil {
		c.JSON(404, gin.H{"error": "user not found"})
		return
	}

	c.JSON(200, user)
}
