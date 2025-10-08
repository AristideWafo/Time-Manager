package middleware

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"testApi/service"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {

		if context.FullPath() == "/authentification" {
			context.Next()
			return
		}

		var token = context.Request.FormValue("api_token")

		if token == "" {
			err := errors.New("NO TOKEN GIVEN")
			context.AbortWithError(http.StatusBadRequest, error(err))
			context.JSON(http.StatusBadRequest, gin.H{"error": "No token given"})
			return
		}

		if service.ValidateToken(token) {
			context.Next()
			return
		}

		err := errors.New("INVALID TOKEN")
		context.AbortWithError(http.StatusBadRequest, error(err))
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token"})
	}
}
