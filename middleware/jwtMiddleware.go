package middleware

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// Secret key for JWT signing
var jwtSecret = []byte("your_secret_key")

// JWTAuthMiddleware is the middleware for securing routes
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("JWTAuthMiddleware Called")
		// Get the token from the Authorization header
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token not provided"})
			c.Abort()
			return
		}

		// Remove the "Bearer " prefix
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		// Parse the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Validate the signing algorithm
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Token is valid, continue processing the request
		c.Next()
	}
}

// GenerateToken generates a JWT token (used in authController)
func GenerateToken(userID int) (string, error) {
	log.Println("GenerateToken Called")
	// Set token expiration time (e.g., 1 hour)
	expirationTime := time.Now().Add(1 * time.Hour)

	// Create the JWT claims
	claims := &jwt.RegisteredClaims{
		Issuer:    fmt.Sprintf("%d", userID), // User ID as issuer
		ExpiresAt: jwt.NewNumericDate(expirationTime),
	}

	// Create the token using the HMAC signing method
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate the signed token string
	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
