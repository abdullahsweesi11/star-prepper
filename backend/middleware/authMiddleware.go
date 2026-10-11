package middleware

import (
	"net/http"
	"star-prepper/models"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var credentials models.AuthModel

		if err := c.ShouldBindJSON(&credentials); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid auth body structure"})
			c.Abort()
			return
		}

		c.Set("credentials", credentials)

		c.Next()
	}
}
