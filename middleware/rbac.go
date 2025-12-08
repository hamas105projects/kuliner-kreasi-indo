package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// allowedRoles contoh: []string{"owner", "cashier"}
func RBACMiddleware(allowedRoles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists {
			c.JSON(http.StatusForbidden, gin.H{"message": "Role not found"})
			c.Abort()
			return
		}

		for _, r := range allowedRoles {
			if role == r {
				c.Next()
				return
			}
		}

		c.JSON(http.StatusForbidden, gin.H{"message": "Access denied"})
		c.Abort()
	}
}
