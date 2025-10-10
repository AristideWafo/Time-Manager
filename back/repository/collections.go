package repository

import "go.mongodb.org/mongo-driver/v2/mongo"

var (
	teamCollection *mongo.Collection
	userCollection *mongo.Collection
)

func TeamCollection() *mongo.Collection {
	if teamCollection == nil {
		teamCollection = Client.Database.Collection("Team")
	}
	return teamCollection
}

func UserCollection() *mongo.Collection {
	if userCollection == nil {
		teamCollection = Client.Database.Collection("User")
	}
	return userCollection
}
