package controllers

import (
	"ecommerce/config"
	"ecommerce/error"
	"ecommerce/middleware"
	"ecommerce/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Login handles authentication and returns a JWT token
func Login(c *gin.Context) {
	log.Println("Login Called")
	var userInput models.User
	if err := c.ShouldBindJSON(&userInput); err != nil {
		models.CreateErrorResponse(c, http.StatusBadRequest, "", error.ErrInvalidRequest)
		return
	}

	// Fetch user from the database
	var user models.User
	if err := config.DB.Where("email = ?", userInput.Email).First(&user).Error; err != nil {
		models.CreateErrorResponse(c, http.StatusUnauthorized, "Invalid Email", error.ErrInvalidCredential)
		return
	}

	// Check if the password is correct
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password)); err != nil {
		models.CreateErrorResponse(c, http.StatusUnauthorized, "Invalid Password", error.ErrInvalidCredential)
		return
	}

	// Generate the JWT token
	token, err := middleware.GenerateToken(user.ID)
	if err != nil {
		models.CreateErrorResponse(c, http.StatusInternalServerError, "Error in generating token", &error.Error{Code: "", Message: "Could not generate token"})
		return
	}
	models.CreateSuccessResponse(c, http.StatusOK, "Login successful", token)
}

func Register(c *gin.Context) {
	log.Println("Register API called")
	var userInput models.User

	if err := c.ShouldBindJSON(&userInput); err != nil {
		models.CreateErrorResponse(c, http.StatusBadRequest, "Validation failed for user name or password please retry", error.ErrInvalidRequest)
		return
	}
	if len(userInput.Password) < 8 {
		models.CreateErrorResponse(c, http.StatusBadRequest, "Password must be at least 8 characters long", error.ErrInvalidRequest)
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	if err != nil {
		models.CreateErrorResponse(c, http.StatusInternalServerError, "Failed to hash password", error.INTERNAL_ERROR)
		return
	}

	// Save the user to the database
	user := models.User{
		Email:    userInput.Email,
		Password: string(hashedPassword),
	}

	if err := config.DB.Create(&user).Error; err != nil {
		models.CreateErrorResponse(c, http.StatusInternalServerError, "Failed to create user", error.INTERNAL_ERROR)
		return
	}

	models.CreateSuccessResponse(c, http.StatusCreated, "User created successfully", nil)
}
