package authentification

type AuthentificationUserInput struct {
	Email    string `json:"email" binding:"required" validate:"required,email"`
	Password string `json:"password" binding:"required" validate:"required"`
}

type AuthentificationUserOutput struct {
	Token     string `json:"token" binding:"required" validate:"required"`
	FirstName string `json:"first_name" binding:"required" validate:"required"`
	LastName  string `json:"last_name" binding:"required" validate:"required"`
	Email     string `json:"email" binding:"required" validate:"required,email"`
	Role      string `json:"role" binding:"required" validate:"required"`
	Team      string `json:"team" binding:"required" validate:"required"`
}
