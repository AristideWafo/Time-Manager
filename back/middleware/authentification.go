package middleware

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"TimeManager/service"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {

		path := context.Request.URL.Path
		allowedPaths := []string{"/swagger/", "/docs/"}
		
		for _, allowedPath := range allowedPaths {
			if len(path) >= len(allowedPath) && path[:len(allowedPath)] == allowedPath {
				context.Next()
				return
			}
		}

		token := context.Request.Header.Get("api_token")

		if token == "" {
			err := errors.New("NO TOKEN GIVEN")
			err = context.AbortWithError(http.StatusBadRequest, err)
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		claims, err := service.ValidateToken(token)

		if err == nil {
			context.Set("claims", claims)
			context.Next()
			return
		}

		err = context.AbortWithError(http.StatusUnauthorized, err)
		context.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})

	}
}
