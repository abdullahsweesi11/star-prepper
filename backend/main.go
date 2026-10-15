package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"star-prepper/controllers"
	"star-prepper/middleware"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	var err error

	// load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Could not load .env file")
	}

	// Open Postgres connection
	db, err = sql.Open(
		"postgres",
		fmt.Sprintf(
			"host=localhost port=5432 user=%s password=%s dbname=%s sslmode=disable",
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("POSTGRES_DB"),
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Verify DB connection is immediately open
	if err := db.Ping(); err != nil {
		db.Close()
		log.Fatal(err)
	}
}

func main() {
	if os.Getenv("APP_ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	defer db.Close()

	router := gin.Default()
	router.Use(middleware.DatabaseMiddleware(db))
	router.Use(middleware.CorsConfig)

	// router.GET("/user", controllers.GetUsers)

	authGroup := router.Group("/auth")
	authGroup.POST("/register", middleware.AuthMiddleware(), controllers.Register)
	authGroup.POST("/login", middleware.AuthMiddleware(), controllers.Login)
	authGroup.POST("/logout", middleware.ProtectedMiddleware(), controllers.Logout)

	authGroup.GET("/me", middleware.ProtectedMiddleware(), controllers.GetCurrentUser)

	router.Run("localhost:8080")
}
