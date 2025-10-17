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
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "TimeManager/docs"
)

//	@title		Time-Manager API
//	@version	0.1

//	@host	localhost:8080

//	@securityDefinitions.basic	BasicAuth

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

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := router.Run("localhost:8080")

	if err != nil {
		panic(err)
	}
}
