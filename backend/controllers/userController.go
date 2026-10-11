package controllers

import (
	"database/sql"
	"net/http"
	"star-prepper/models"

	"github.com/gin-gonic/gin"
)

func GetCurrentUser(c *gin.Context) {
	user := c.MustGet("currentUser").(models.User)

	c.JSON(http.StatusOK, gin.H{
		"user": map[string]string{"email": user.Email},
	})
}

func GetUsers(c *gin.Context) {
	db, ok := c.MustGet("db").(*sql.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not access DB (2)"})
	}

	rows, err := db.QueryContext(c.Request.Context(), "SELECT email FROM users")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
		return
	}
	defer rows.Close()

	var users []models.UserPublic

	for rows.Next() {
		var u models.UserPublic
		if err := rows.Scan(&u.Email); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not parse row"})
			return
		}
		users = append(users, u)
	}

	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error while iterating over rows"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}
