package main

import (
	"testApi/middleware"
	"testApi/route/autentification"
	"testApi/route/user"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.Use(middleware.AuthMiddleware())

	autentification.RegisterAuthentificationRoutes(router)
	user.RegisterUserRoutes(router)

	router.Run("localhost:8080")
}
