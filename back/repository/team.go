package repository

import (
	"TimeManager/model"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func FetchTeam(name string) (model.Team, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var result model.Team
	filter := bson.M{"name": name}
	err := TeamCollection().FindOne(ctx, filter).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Printf("No team found: %v\n", err)
		} else {
			fmt.Printf("Error found : %v", err)
		}
		return model.Team{}, false, err
	}

	err = model.ValidateModel(&result)

	if err != nil {
		fmt.Printf("Error found : %v", err)
		return model.Team{}, false, err
	}

	fmt.Printf("%v", result)
	return result, true, nil
}

func FetchAllTeams() ([]model.Team, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var result []model.Team

	cursor, err := TeamCollection().Find(ctx, bson.M{})
	if err != nil {
		return []model.Team{}, false, fmt.Errorf("failed to find teams: %w", err)
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &result); err != nil {
		return []model.Team{}, false, fmt.Errorf("failed to decode teams: %w", err)
	}

	return result, true, nil
}

func CreateTeam(team model.Team) (model.Team, bool, error) {
	return team, true, nil
}

func UpdateTeam(team model.Team) (model.Team, bool, error) {
	return team, true, nil
}

func DeleteTeam(team model.Team) (model.Team, bool, error) {
	return team, true, nil
}
