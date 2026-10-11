package controllers

import (
	"database/sql"
	"encoding/hex"
	"net/http"
	"star-prepper/models"

	"crypto/sha256"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	// get DB connection
	db, ok := c.MustGet("db").(*sql.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not access DB"})
		return
	}

	// get inputted credentials
	credentials, ok := c.MustGet("credentials").(models.AuthModel)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not access credentials"})
		return
	}

	// check whether user already exists
	var exists bool
	query := "SELECT EXISTS(SELECT email FROM users WHERE email = $1)"
	err := db.QueryRowContext(c.Request.Context(), query, credentials.Email).Scan(&exists)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not access database"})
		return
	}
	if exists {
		c.JSON(http.StatusConflict, gin.H{"error": "user already exists"})
		return
	}

	// fits password into 32-byte pre-hash (as hex string to prevent null bytes)
	checksum := sha256.Sum256([]byte(credentials.Password))
	preHash := hex.EncodeToString(checksum[:])

	// pass pre-hash to main bcrypt hashing algorithm
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(preHash), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not register user"})
		return
	}

	// insert into DB
	query = "INSERT INTO users (email, hashed_password) VALUES ($1, $2)"
	_, err = db.ExecContext(c.Request.Context(), query, credentials.Email, hashedPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not register user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "created user"})
}

// func Login(c *gin.Context) {

// }
