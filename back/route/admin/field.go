package admin

type CreateUserInput struct {
	FirstName string `json:"FirstName" bson:"FirstName" binding:"required" validate:"required"`
	LastName  string `json:"LastName" bson:"LastName" binding:"required" validate:"required"`
	Role      string `json:"Role" binding:"required" validate:"required"`
	Email     string `json:"Email" binding:"required" validate:"required,email"`
	Password  string `json:"Password" binding:"required" validate:"required"`
	Team      string `json:"Team"`
}

type UpdateUserTeamInput struct {
	TeamId string `json:"TeamId" bson:"TeamId" binding:"required" validate:"required"`
}

type CreateTeamInput struct {
	Name string `json:"Name" bson:"Name" binding:"required" validate:"required"`
}

type UpdateTeamInput struct {
	CurrentName string `json:"CurrentName" bson:"CurrentName" binding:"required" validate:"required"`
	NewName     string `json:"NewName" bson:"NewName" binding:"required" validate:"required"`
}

type RequestedTeamInput struct {
	Name string `json:"Name" bson:"Name" binding:"required" validate:"required"`
}
