package controllers

import (
	"database/sql"
	"encoding/hex"
	"net/http"
	"os"
	"star-prepper/models"
	"star-prepper/utils"
	"time"

	"crypto/sha256"

	"github.com/gin-gonic/gin"
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
	var userId string
	query = "INSERT INTO users (email, hashed_password) VALUES ($1, $2) RETURNING user_id"
	err = db.QueryRowContext(c.Request.Context(), query, credentials.Email, hashedPassword).Scan(&userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not register user"})
		return
	}

	// user authenticated

	setAuthCookieAndResponse(c, http.StatusCreated, userId, credentials.Email, "registration (and login) successful")
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
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not access database"})
		return
	}

	checksum := sha256.Sum256([]byte(credentials.Password))
	preHash := hex.EncodeToString(checksum[:])

	if err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(preHash)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	// user has been authenticated

	setAuthCookieAndResponse(c, http.StatusOK, user.UserId.String(), user.Email, "login sucessful")
}

func Logout(c *gin.Context) {
	c.SetCookie("access_token", "", -1, "/",
		"localhost", // TODO: change to domain name in prod
		false,       // TODO: change to true in prod
		true,
	)

	c.JSON(http.StatusOK, gin.H{"message": "logout successful"})
}

func setAuthCookieAndResponse(c *gin.Context, statusCode int, userId string, email string, message string) {
	durationHours := time.Hour * 24

	token, err := utils.GenerateJWT(userId, time.Now().Add(durationHours).Unix())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not sign JWT token"})
		return
	}

	isProd := gin.Mode() == gin.ReleaseMode
	domain := os.Getenv("COOKIE_DOMAIN")

	c.SetCookie(
		"access_token", token,
		int(durationHours.Seconds()),
		"/",
		domain,
		isProd,
		true,
	)

	c.JSON(statusCode, gin.H{
		"message": message,
		"user":    map[string]string{"userId": userId, "email": email},
	})
}
