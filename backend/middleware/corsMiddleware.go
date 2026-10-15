package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
)

var CorsConfig = cors.New(cors.Config{
	AllowOrigins:     []string{"http://127.0.0.1:5173", "http://localhost:5173"},
	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
	ExposeHeaders:    []string{"Content-Length"},
	AllowCredentials: true,
	MaxAge:           12 * time.Hour,
})
