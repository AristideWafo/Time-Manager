package main

import (
	"testApi/middleware"
	"testApi/route/autentification"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.Use(middleware.AuthMiddleware())

	autentification.RegisterAuthentificationRoutes(router)

	router.Run("localhost:8080")
}
