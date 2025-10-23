package repository

import "go.mongodb.org/mongo-driver/v2/mongo"

var (
	teamCollection     *mongo.Collection
	userCollection     *mongo.Collection
	presenceCollection *mongo.Collection
)

func TeamCollection() *mongo.Collection {
	if teamCollection == nil {
		teamCollection = Client.Database.Collection("Team")
	}
	return teamCollection
}

func UserCollection() *mongo.Collection {
	if userCollection == nil {
		userCollection = Client.Database.Collection("User")
	}
	return userCollection
}

func PresenceCollection() *mongo.Collection {
	if presenceCollection == nil {
		presenceCollection = Client.Database.Collection("Presence")
	}
	return presenceCollection
}
