package repository

import (
	"context"
	"os"
	"testing"

	"TimeManager/model"

	"github.com/subosito/gotenv"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var testCollectionName = "test_collection"

func setupTestConnection(t *testing.T) {
	err := gotenv.Load("../.env")
	if err != nil {
		t.Fatalf("failed to load .env file: %v", err)
	}

	os.Setenv("DB_NAME", "test")

	DBConnect()
}

func setupTestCollection(t *testing.T) *mongo.Collection {

	setupTestConnection(t)

	if Client.Database == nil {
		t.Skip("Database not connected")
	}
	return Client.Database.Collection(testCollectionName)
}

func cleanupTestCollection(_ *testing.T) {
	if Client.Database != nil {
		Client.Database.Collection(testCollectionName).DeleteMany(context.TODO(), bson.D{})
	}
}

func TestSave_ValidDocument(t *testing.T) {
	collection := setupTestCollection(t)
	defer cleanupTestCollection(t)

	user := &model.User{
		Password:  "hashed_password",
		FirstName: "Doctor",
		LastName:  "Who",
		Email:     "Doctor@Tardis.com",
		Role:      "EMPLOYEE",
	}

	err := Save(user, collection)
	if err != nil {
		t.Fatalf("expected no error for valid save, got %v", err)
	}

	if user.ID.IsZero() {
		t.Fatalf("expected ID to be set after save, got zero")
	}

	if user.CreatedAt.IsZero() {
		t.Fatalf("expected CreatedAt to be set after save")
	}

	if user.UpdatedAt.IsZero() {
		t.Fatalf("expected UpdatedAt to be set after save")
	}
}

func TestSave_InvalidDocument(t *testing.T) {
	collection := setupTestCollection(t)
	defer cleanupTestCollection(t)

	user := &model.User{
		Password: "password",
		Email:    "Doctor@Tardis.com",
		Role:     "EMPLOYEE",
	}

	err := Save(user, collection)
	if err == nil {
		t.Fatalf("expected error for invalid document (missing FirstName), got nil")
	}
}

func TestUpdateOne_ValidUpdate(t *testing.T) {
	collection := setupTestCollection(t)
	defer cleanupTestCollection(t)

	user := &model.User{
		Password:  "hashed_password",
		FirstName: "Doctor",
		LastName:  "Who",
		Email:     "Doctor@Tardis.com",
		Role:      "EMPLOYEE",
	}

	err := Save(user, collection)
	if err != nil {
		t.Fatalf("failed to save user: %v", err)
	}

	originalID := user.ID
	user.FirstName = "River"

	err = UpdateOne(user, collection, bson.D{{Key: "FirstName", Value: "River"}}, bson.D{})
	if err != nil {
		t.Fatalf("expected no error for valid update, got %v", err)
	}

	if user.ID != originalID {
		t.Fatalf("expected ID to remain unchanged")
	}

	if user.FirstName != "River" {
		t.Fatalf("expected FirstName to be updated to River, got %s", user.FirstName)
	}
}

func TestUpdateOne_InvalidFieldUpdate(t *testing.T) {
	collection := setupTestCollection(t)
	defer cleanupTestCollection(t)

	user := &model.User{
		Password:  "hashed_password",
		FirstName: "Doctor",
		LastName:  "Who",
		Email:     "Doctor@Tardis.com",
		Role:      "EMPLOYEE",
	}

	err := Save(user, collection)
	if err != nil {
		t.Fatalf("failed to save user: %v", err)
	}

	err = UpdateOne(user, collection, bson.D{{Key: "NonExistentField", Value: "value"}}, bson.D{})
	if err == nil {
		t.Fatalf("expected error for invalid field update, got nil")
	}
}

func TestGetOne_DocumentExists(t *testing.T) {
	collection := setupTestCollection(t)
	defer cleanupTestCollection(t)

	user := &model.User{
		Password:  "hashed_password",
		FirstName: "Doctor",
		LastName:  "Who",
		Email:     "Doctor@Tardis.com",
		Role:      "EMPLOYEE",
	}

	err := Save(user, collection)
	if err != nil {
		t.Fatalf("failed to save user: %v", err)
	}

	savedID := user.ID

	retrieved := &model.User{}
	filter := bson.D{{Key: "_id", Value: savedID}}
	err = GetOne(filter, collection, retrieved)
	if err != nil {
		t.Fatalf("expected no error for existing document, got %v", err)
	}

	if retrieved.ID != savedID {
		t.Fatalf("expected retrieved ID to match saved ID")
	}

	if retrieved.FirstName != "Doctor" {
		t.Fatalf("expected retrieved FirstName to be John, got %s", retrieved.FirstName)
	}
}

func TestGetOne_DocumentNotExists(t *testing.T) {
	collection := setupTestCollection(t)
	defer cleanupTestCollection(t)

	nonExistentID, _ := bson.ObjectIDFromHex("507f1f77bcf86cd799439999")
	retrieved := &model.User{}
	filter := bson.D{{Key: "_id", Value: nonExistentID}}
	err := GetOne(filter, collection, retrieved)
	if err == nil {
		t.Fatalf("expected error for non-existent document, got nil")
	}
}

func TestGetMany_MultipleDocuments(t *testing.T) {
	collection := setupTestCollection(t)
	defer cleanupTestCollection(t)

	user1 := &model.User{
		Password:  "hashed_password",
		FirstName: "Doctor",
		LastName:  "Who",
		Email:     "Doctor@Tardis.com",
		Role:      "EMPLOYEE",
	}

	user2 := &model.User{
		Password:  "hashed_password",
		FirstName: "River",
		LastName:  "Smith",
		Email:     "jane@Tardis.com",
		Role:      "MANAGER",
	}

	err := Save(user1, collection)
	if err != nil {
		t.Fatalf("failed to save user1: %v", err)
	}

	err = Save(user2, collection)
	if err != nil {
		t.Fatalf("failed to save user2: %v", err)
	}

	retrieved := make([]*model.User, 0)
	filter := bson.D{}
	err = GetMany(filter, collection, &retrieved)
	if err != nil {
		t.Fatalf("expected no error for GetMany, got %v", err)
	}

	if len(retrieved) < 2 {
		t.Fatalf("expected at least 2 documents, got %d", len(retrieved))
	}
}

func TestGetMany_EmptyResult(t *testing.T) {
	collection := setupTestCollection(t)
	defer cleanupTestCollection(t)

	nonExistentID, _ := bson.ObjectIDFromHex("507f1f77bcf86cd799439999")
	retrieved := make([]*model.User, 0)
	filter := bson.D{{Key: "_id", Value: nonExistentID}}
	err := GetMany(filter, collection, &retrieved)
	if err != nil {
		t.Fatalf("expected no error for empty GetMany, got %v", err)
	}

	if len(retrieved) != 0 {
		t.Fatalf("expected 0 documents, got %d", len(retrieved))
	}
}

func TestDeleteOne_DocumentExists(t *testing.T) {
	collection := setupTestCollection(t)
	defer cleanupTestCollection(t)

	user := &model.User{
		Password:  "hashed_password",
		FirstName: "Doctor",
		LastName:  "Who",
		Email:     "Doctor@Tardis.com",
		Role:      "EMPLOYEE",
	}

	err := Save(user, collection)
	if err != nil {
		t.Fatalf("failed to save user: %v", err)
	}

	savedID := user.ID

	err = DeleteOne(user, collection)
	if err != nil {
		t.Fatalf("expected no error for delete, got %v", err)
	}

	retrieved := &model.User{}
	filter := bson.D{{Key: "_id", Value: savedID}}
	err = GetOne(filter, collection, retrieved)
	if err == nil {
		t.Fatalf("expected error after delete, got nil")
	}
}

func TestUpdateMany_MultipleDocuments(t *testing.T) {
	collection := setupTestCollection(t)
	defer cleanupTestCollection(t)

	user1 := &model.User{
		Password:  "hashed_password",
		FirstName: "Doctor",
		LastName:  "Who",
		Email:     "Doctor@Tardis.com",
		Role:      "EMPLOYEE",
	}

	user2 := &model.User{
		Password:  "hashed_password",
		FirstName: "River",
		LastName:  "Who",
		Email:     "jane@Tardis.com",
		Role:      "EMPLOYEE",
	}

	err := Save(user1, collection)
	if err != nil {
		t.Fatalf("failed to save user1: %v", err)
	}

	err = Save(user2, collection)
	if err != nil {
		t.Fatalf("failed to save user2: %v", err)
	}

	users := []*model.User{user1, user2}
	filter := bson.D{{Key: "LastName", Value: "Who"}}

	err = UpdateMany(&users, filter, collection, bson.D{{Key: "FirstName", Value: "Updated"}}, bson.D{})
	if err != nil {
		t.Fatalf("expected no error for UpdateMany, got %v", err)
	}

	if users[0].FirstName != "Updated" {
		t.Fatalf("expected user1 FirstName to be Updated, got %s", users[0].FirstName)
	}

	if users[1].FirstName != "Updated" {
		t.Fatalf("expected user2 FirstName to be Updated, got %s", users[1].FirstName)
	}
}

func TestUpdateMany_EmptySlice(t *testing.T) {
	collection := setupTestCollection(t)
	defer cleanupTestCollection(t)

	users := make([]*model.User, 0)
	filter := bson.D{}

	err := UpdateMany(&users, filter, collection, bson.D{}, bson.D{})
	if err != nil {
		t.Fatalf("expected no error for empty slice, got %v", err)
	}
}
