package middleware

import (
	"httpServer/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(authService *services.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get token from header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header not found"})
			c.Abort()
			return
		}

		// Get token from header
		token := strings.Split(authHeader, " ")
		if len(token) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header"})
			c.Abort()
			return
		}

		// Verify token
		userID, err := authService.VerifyAccessJWT(token[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Set user ID to context
		c.Set("user_id", userID)
		c.Next()
	}
}
