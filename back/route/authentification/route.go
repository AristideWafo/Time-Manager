package authentification

import (
	"TimeManager/model"
	"TimeManager/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func RegisterAuthentificationRoutes(router *gin.Engine) {
	auth := router.Group("/api/authentification")
	{
		auth.POST("", authentificate)
	}
}

// authentificate
//
//	@Summary		Authenticate a user
//	@Description	Authenticates a user using email and password, and returns a JWT token if successful
//	@Tags			Authentication
//	@Accept			json
//	@Produce		json
//	@Param			user	body		AuthentificationUserInput	true	"User credentials"
//	@Success		200		{object}	AuthentificationUserOutput
//	@Failure		400		"Invalid input"
//	@Failure		404		"User not found or wrong credentials"
//	@Failure		500		"Internal server error"
//	@Router			/api/authentification [post]
func authentificate(context *gin.Context) {

	var user AuthentificationUserInput

	if err := context.BindJSON(&user); err != nil {
		err = context.AbortWithError(http.StatusBadRequest, err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fetched_user, err := service.Login(user.Email, user.Password)

	if err != nil {
		err = context.AbortWithError(http.StatusNotFound, err)
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	token, err := service.CreateToken(fetched_user.ID, fetched_user.Role, fetched_user.Team)

	if err != nil {
		err = context.AbortWithError(http.StatusInternalServerError, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var output AuthentificationUserOutput

	if err = copier.Copy(&output, &fetched_user); err != nil {
		err = context.AbortWithError(http.StatusInternalServerError, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	output.Token = token
	output.ID = fetched_user.ID.Hex()

	if fetched_user.Team != bson.NilObjectID {
		output.Team = fetched_user.Team.Hex()
	}

	if err = model.ValidateModel(&output); err != nil {
		err = context.AbortWithError(http.StatusInternalServerError, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, output)

}
