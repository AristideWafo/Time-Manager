package middleware

import (
	"TimeManager/service"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BasePrivilegesMiddleWare(context *gin.Context, privileges []string, concerned_routes []string) {

	unverified := true

	for _, route := range concerned_routes {
		if context.FullPath() == route {
			unverified = false
			break
		}
	}

	if unverified {
		context.Next()
		return
	}

	claims, ok := context.Get("claims")

	if !ok {
		err := errors.New("ERROR WHEN GETTING CLAIMS")
		err = context.AbortWithError(http.StatusInternalServerError, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	asserted_claims, ok := claims.(service.TokenClaims)

	if !ok {
		err := errors.New("ERROR WHEN GETTING CLAIMS")
		err = context.AbortWithError(http.StatusInternalServerError, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for _, privilege := range privileges {
		if asserted_claims.Audience[0] == privilege {
			context.Next()
			return
		}
	}

	err_msg := fmt.Sprintf("YOU NEED %v PRIVILEGES", privileges)
	err := errors.New(err_msg)
	err = context.AbortWithError(http.StatusBadRequest, err)
	context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

}

func ManagerMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		concerned_routes := []string{}
		BasePrivilegesMiddleWare(context, []string{"MANAGER", "ADMIN"}, concerned_routes)
	}
}

func AdminMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		concerned_routes := []string{"/api/user/create"}
		BasePrivilegesMiddleWare(context, []string{"ADMIN"}, concerned_routes)
	}
}
