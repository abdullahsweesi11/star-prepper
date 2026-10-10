package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"star-prepper/models"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	fmt.Println("hi")
	db := c.MustGet("db").(*sql.DB)

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
