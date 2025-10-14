package user

type CreateUserInput struct {
	FirstName string `json:"FirstName" bson:"FirstName" binding:"required" validate:"required"`
	LastName  string `json:"LastName" bson:"LastName" binding:"required" validate:"required"`
	Role      string `json:"Role" binding:"required" validate:"required"`
	Email     string `json:"Email" binding:"required" validate:"required,email"`
	Password  string `json:"Password" binding:"required" validate:"required"`
	Team      string `json:"Team"`
}

type UserOutput struct {
	ID        string `json:"_id" bson:"_id" binding:"required" validate:"required"`
	FirstName string `json:"FirstName" bson:"FirstName" binding:"required" validate:"required"`
	LastName  string `json:"LastName" bson:"LastName" binding:"required" validate:"required"`
	Email     string `json:"Email" binding:"required" validate:"required,email"`
	Role      string `json:"Role" binding:"required" validate:"required"`
	Team      string `json:"Team"`
}
