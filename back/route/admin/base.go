package admin

import (
	"TimeManager/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterAdminRoutes(router *gin.Engine) {
	admin := router.Group("/api/admin").Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
	{
		admin.GET("/:id", GetUserById)
		admin.POST("/user/create", PostUser)
		admin.DELETE("/user/delete/:id", DeleteUserById)
		admin.PUT("/update/team/user/:id", UpdateUserTeam)
		admin.GET("/all/user", GetAllUsers)
		admin.GET("/all/team", GetAllTeams)
		admin.GET("/team/users/:name", GetAllUsersFromTeam)
		admin.POST("/team/create", PostTeam)
		admin.PUT("/team/update", UpdateTeam)
		admin.DELETE("/team/delete/:id", DeleteTeamById)
	}
}
