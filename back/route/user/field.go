package user

type UserInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserOutput struct {
	Username string `json:"username" binding:"required"`
	Role     string `json:"role" binding:"required"`
	Team     string `json:"team" binding:"required"`
}
