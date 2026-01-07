package middleware

import (
	"TimeManager/service"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BasePrivilegesMiddleWare(context *gin.Context, privileges []string) bool {

	claims, ok := context.Get("claims")

	if !ok {
		err := errors.New("ERROR WHEN GETTING CLAIMS")
		err = context.AbortWithError(http.StatusBadRequest, err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return false
	}

	asserted_claims, ok := claims.(service.TokenClaims)

	if !ok {
		err := errors.New("ERROR WHEN GETTING CLAIMS")
		err = context.AbortWithError(http.StatusBadRequest, err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return false
	}

	for _, privilege := range privileges {
		if asserted_claims.Audience[0] == privilege {
			return true
		}
	}

	err_msg := fmt.Sprintf("YOU NEED %v PRIVILEGES", privileges)
	err := errors.New(err_msg)
	err = context.AbortWithError(http.StatusBadRequest, err)
	context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

	return false

}

func ManagerMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		if !BasePrivilegesMiddleWare(context, []string{"MANAGER", "ADMIN"}) {
			return
		}
		context.Next()
	}
}

func AdminMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		if !BasePrivilegesMiddleWare(context, []string{"ADMIN"}) {
			return
		}
		context.Next()
	}
}
