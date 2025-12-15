package teams

import (
	"TimeManager/middleware"
	"TimeManager/model"
	"TimeManager/service"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

func RegisterTeamRoutes(router *gin.Engine) {
	team := router.Group("/api/team").Use(middleware.ManagerMiddleware())
	{
		team.GET("/:name", FetchTeam)
	}
}

// FetchTeam
//
// @Summary		Fetch a single team
// @Description	Fetch a single team and returns its info if successful
// @Tags		Team
// @Produce		json
// @Success		200		{object}	TeamOutput
// @Failure		400		"Invalid Input"
// @Failure		404		"Team not found"
// @Failure		500		"Internal server error"
// @Router		/api/team/name [get]
// @Security 	ApiKeyAuth
func FetchTeam(context *gin.Context) {
	name := context.Param("name")

	if name == "" {
		err := errors.New("team name must be specified")
		err = context.AbortWithError(http.StatusBadRequest, err)
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	team, err := service.GetTeamByName(name)

	if err != nil {
		err = context.AbortWithError(http.StatusNotFound, err)
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var output TeamOutput

	if err = copier.Copy(&output, team); err != nil {
		err = context.AbortWithError(http.StatusInternalServerError, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	output.ID = team.ID.Hex()

	if err = model.ValidateModel(&output); err != nil {
		err = context.AbortWithError(http.StatusInternalServerError, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"team": output})

}
