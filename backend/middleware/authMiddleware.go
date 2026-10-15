package middleware

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"star-prepper/models"

	"github.com/gin-gonic/gin"
	jwt5 "github.com/golang-jwt/jwt/v5"
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

func ProtectedMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("access_token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user must be authorized"})
			c.Abort()
			return
		}

		token, err := jwt5.Parse(tokenString, func(token *jwt5.Token) (any, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		}, jwt5.WithValidMethods([]string{"HS256"}))

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt5.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		// get DB connection
		db, ok := c.MustGet("db").(*sql.DB)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not access DB (3)"})
			c.Abort()
			return
		}

		userId, ok := claims["id"].(string)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token claims"})
			c.Abort()
			return
		}

		var user models.User
		query := "SELECT user_id, email, hashed_password, role_id, created_at FROM users WHERE user_id = $1"

		err = db.QueryRowContext(
			c.Request.Context(), query, userId,
		).Scan(
			&user.UserId, &user.Email, &user.HashedPassword, &user.RoleId, &user.CreatedAt,
		)
		if err == sql.ErrNoRows {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user does not exist"})
			c.Abort()
			return
		}
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not access database"})
			c.Abort()
			return
		}

		c.Set("currentUser", user)

		c.Next()
	}
}
