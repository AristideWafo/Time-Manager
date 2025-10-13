package main

import (
	"TimeManager/middleware"
	"TimeManager/repository"
	"TimeManager/route/authentification"
	"TimeManager/route/presence"
	"TimeManager/route/user"

	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"
)

func init() {
	err := gotenv.Load("../.env")
	if err != nil {
		panic(err)
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
