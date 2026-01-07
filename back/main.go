package main

import (
	"TimeManager/repository"
	"TimeManager/route/admin"
	"TimeManager/route/authentification"
	"TimeManager/route/manager"
	"TimeManager/route/public/presence"
	teams "TimeManager/route/public/team"
	users "TimeManager/route/public/user"

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

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name api_token

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

	authentification.RegisterAuthentificationRoutes(router)
	users.RegisterUserRoutes(router)
	teams.RegisterTeamRoutes(router)
	presence.RegisterPresenceRoutes(router)
	admin.RegisterAdminRoutes(router)
	manager.RegisterManagerRoutes(router)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := router.Run("localhost:8080")

	if err != nil {
		panic(err)
	}
}
