package middleware

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"TimeManager/service"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {

		if context.FullPath() == "/authentification" {
			context.Next()
			return
		}

		token := context.Request.Header.Get("Api_token")

		if token == "" {
			err := errors.New("NO TOKEN GIVEN")
			context.AbortWithError(http.StatusBadRequest, err)
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		claims, err := service.ValidateToken(token)

		if err == nil {
			context.Set("claims", claims)
			context.Next()
			return
		}

		context.AbortWithError(http.StatusBadRequest, err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

	}
}
