package middleware

import (
	"context"
	"fmt"
	"httpServer/config"
	"httpServer/services"
	"net/http"
	"strconv"
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
		token := strings.TrimPrefix(authHeader, "Bearer ")

		// parse token and get username
		userID, err := services.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		// check token in redis
		accessToken, err := config.RedisAccessTokenClient.Get(context.Background(), userID).Result()
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token not found or expired"})
			c.Abort()
			return
		}

		// compare token
		if accessToken != token {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token mismatch"})
			c.Abort()
			return
		}

		userIDUint, err := strconv.ParseUint(userID, 10, 32)
		if err != nil {
			fmt.Println("Lỗi:", err)
			return
		}

		// Set user ID to context
		c.Set("user_id", uint(userIDUint))
		c.Next()
	}
}
