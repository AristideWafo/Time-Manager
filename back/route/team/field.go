package team

type TeamOutput struct {
	ID   string `json:"_id" bson:"_id" binding:"required" validate:"required"`
	Name string `json:"Name" bson:"Name" binding:"required" validate:"required"`
}

type CreateTeamInput struct {
	Name string `json:"Name" bson:"Name" binding:"required" validate:"required"`
}

type UpdateTeamInput struct {
	CurrentName string `json:"CurrentName" bson:"CurrentName" binding:"required" validate:"required"`
	NewName     string `json:"NewName" bson:"NewName" binding:"required" validate:"required"`
}
