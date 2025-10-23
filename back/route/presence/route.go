package presence

import (
	"TimeManager/model"
	"TimeManager/service"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func RegisterPresenceRoutes(router *gin.Engine) {
	user := router.Group("/api/presence")
	{
		user.GET("", GetAllPresences)
		user.POST("/create", CreatePresence)
	}
}

// GetAllPresences godoc
//
//	@Summary		Return all presences for the authenticated user
//	@Description	Returns all presences belonging to the user identified by the JWT token.
//	@Tags			Presence
//	@Accept			json
//	@Produce		json
//	@Success		200		{array}	PresenceOutput
//	@Failure		404		"User not found"
//	@Failure		500		"Internal server error"
//	@Router			/api/presence [get]
func GetAllPresences(context *gin.Context) {

	claims, exists := context.Get("claims")

	if !exists {
		err := errors.New("INTERNAL ISSUE WITH TOKEN")
		err = context.AbortWithError(http.StatusInternalServerError, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	asserted_claims, ok := claims.(service.TokenClaims)

	if !ok {
		err := errors.New("INTERNAL ISSUE WITH TOKEN CONTENT")
		err = context.AbortWithError(http.StatusInternalServerError, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	_id, err := bson.ObjectIDFromHex(asserted_claims.Subject)

	if err != nil {
		err = context.AbortWithError(http.StatusInternalServerError, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	presences, err := service.GetAllPresencesByUserID(_id)

	if err != nil {
		err = context.AbortWithError(http.StatusNotFound, err)
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var output []PresenceOutput

	if err = copier.Copy(&output, presences); err != nil {
		err = context.AbortWithError(http.StatusInternalServerError, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for _, presence := range output {

		if err = model.ValidateModel(&presence); err != nil {
			err = context.AbortWithError(http.StatusInternalServerError, err)
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	context.JSON(http.StatusOK, gin.H{"user": output})
}

// CreatePresence godoc
//
//	@Summary		Create a new presence record
//	@Description	Create a new presence entry for the authenticated user. Returns conflict if duplicate presence exists.// @Tags         Presence
//	@Accept			json
//	@Produce		json
//	@Param			presence	body		CreatePresenceInput	true	"Presence to create"
//	@Success		200			{object}	PresenceOutput
//	@Failure		400			"Invalid input"
//	@Failure		409			"Conflict : presence with same information already exists"
//	@Failure		500			"Internal server error"
//	@Router			/api/presence/create [post]
func CreatePresence(context *gin.Context) {

	claims, exists := context.Get("claims")

	if !exists {
		err := errors.New("INTERNAL ISSUE WITH TOKEN")
		err = context.AbortWithError(http.StatusInternalServerError, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	asserted_claims, ok := claims.(service.TokenClaims)

	if !ok {
		err := errors.New("INTERNAL ISSUE WITH TOKEN CONTENT")
		err = context.AbortWithError(http.StatusInternalServerError, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	_id, err := bson.ObjectIDFromHex(asserted_claims.Subject)

	if err != nil {
		err = context.AbortWithError(http.StatusInternalServerError, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var presence CreatePresenceInput

	if err := context.BindJSON(&presence); err != nil {
		err = context.AbortWithError(http.StatusBadRequest, err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created_presence, err := service.CreatePresence(presence.Type, _id, presence.Timestamp)

	if err != nil {
		err = context.AbortWithError(http.StatusConflict, err)
		context.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	var output PresenceOutput

	if err = copier.Copy(&output, created_presence); err != nil {
		err = context.AbortWithError(http.StatusInternalServerError, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"presence": output})
}
