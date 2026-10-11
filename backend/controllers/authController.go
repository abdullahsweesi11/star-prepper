package controllers

import (
	"database/sql"
	"encoding/hex"
	"fmt"
	"net/http"
	"os"
	"star-prepper/models"
	"time"

	"crypto/sha256"

	"github.com/gin-gonic/gin"
	jwt5 "github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	// get DB connection
	db, ok := c.MustGet("db").(*sql.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not access DB (0)"})
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

func Login(c *gin.Context) {

	// get DB connection
	db, ok := c.MustGet("db").(*sql.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not access DB (1)"})
		return
	}

	// get inputted credentials
	credentials, ok := c.MustGet("credentials").(models.AuthModel)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not access credentials"})
		return
	}

	// check whether user exists
	var user models.UserSecret
	query := "SELECT user_id, email, hashed_password FROM users WHERE email = $1"
	err := db.QueryRowContext(c.Request.Context(), query, credentials.Email).Scan(&user.UserId, &user.Email, &user.HashedPassword)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid credentials"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not access database"})
		return
	}

	checksum := sha256.Sum256([]byte(credentials.Password))
	preHash := hex.EncodeToString(checksum[:])

	if err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(preHash)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid credentials"})
		return
	}

	generateToken := jwt5.NewWithClaims(jwt5.SigningMethodHS256, jwt5.MapClaims{
		"id":  user.UserId.String(),
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := generateToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not sign JWT token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
