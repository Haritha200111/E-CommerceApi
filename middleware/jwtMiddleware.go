package middleware

import (
	"github.com/gin-gonic/gin"
)

// Secret key for JWT signing
var jwtSecret = []byte("your_secret_key")

// JWTAuthMiddleware is the middleware for securing routes
func JWTAuthMiddleware() gin.HandlerFunc {
	// [JWT middleware code here]
}

// GenerateToken generates a token (used in authController)
func GenerateToken(userID uint) (string, error) {
	// [Token generation code here]
}
