package main

import (
	"TimeManager/middleware"
	"TimeManager/repository"
	"TimeManager/route/autentification"
	"TimeManager/route/user"

	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load("../.env")
}

func main() {

	repository.DBConnect()

	router := gin.Default()
	router.Use(middleware.AuthMiddleware())

	autentification.RegisterAuthentificationRoutes(router)
	user.RegisterUserRoutes(router)

	router.Run("localhost:8080")
}
