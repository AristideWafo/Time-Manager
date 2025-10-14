package main

import (
	"TimeManager/middleware"
	"TimeManager/repository"
	"TimeManager/route/authentification"
	"TimeManager/route/presence"
	"TimeManager/route/user"

	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"
)

func init() {
	if os.Getenv("ENV") != "production" && os.Getenv("ENV") != "docker" {
		err := gotenv.Load(".env")
		if err != nil {
			log.Println("Warning: .env file not found, using environment variables")
		}
	}
}

func main() {

	repository.DBConnect()

	router := gin.Default()
	router.Use(middleware.AuthMiddleware())

	authentification.RegisterAuthentificationRoutes(router)
	user.RegisterUserRoutes(router)
	presence.RegisterPresenceRoutes(router)

	err := router.Run("localhost:8080")

	if err != nil {
		panic(err)
	}
}
