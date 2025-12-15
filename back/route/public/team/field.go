package teams

type TeamOutput struct {
	ID   string `json:"_id" bson:"_id" binding:"required" validate:"required"`
	Name string `json:"Name" bson:"Name" binding:"required" validate:"required"`
}
