package main

import (
	"TimeManager/middleware"
	"TimeManager/route/autentification"
	"TimeManager/route/user"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.Use(middleware.AuthMiddleware())

	autentification.RegisterAuthentificationRoutes(router)
	user.RegisterUserRoutes(router)

	router.Run("localhost:8080")
}
