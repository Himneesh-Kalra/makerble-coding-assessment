package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	// "github.com/Himneesh-Kalra/makerble-coding-assessment/models"
	"github.com/Himneesh-Kalra/makerble-coding-assessment/utils"
)

// AuthMiddleware checks JWT and optionally enforces allowed roles.
func AuthMiddleware(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid Authorization header"})
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		user, err := utils.ParseAndVerifyToken(tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		}

		// Optional: check if user's role is allowed
		if len(allowedRoles) > 0 {
			allowed := false
			for _, role := range allowedRoles {
				if user.Role == role {
					allowed = true
					break
				}
			}
			if !allowed {
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Access denied"})
				return
			}
		}

		// Inject user into context
		c.Set("user", user)

		c.Next()
	}
}
