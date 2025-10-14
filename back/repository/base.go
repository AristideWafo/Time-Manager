package repository

import (
	"context"
	"errors"
	"fmt"
	"os"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

var NULL_ID, _ = bson.ObjectIDFromHex("")

type DbClient struct {
	Client   *mongo.Client
	Database *mongo.Database
}

var Client DbClient = DbClient{}

func DBConnect() {

	var DB_USERNAME string = os.Getenv("DB_USERNAME")
	var DB_PASSWORD string = os.Getenv("DB_PASSWORD")
	var DB_HOST string = os.Getenv("DB_HOST")
	var DB_NAME string = os.Getenv("DB_NAME")

	// Use the SetServerAPIOptions() method to set the version of the Stable API on the client
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	var URI string = "mongodb+srv://" + DB_USERNAME + ":" + DB_PASSWORD + "@" + DB_HOST + "/?retryWrites=true&w=majority&appName=Cluster0"
	opts := options.Client().ApplyURI(URI).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(opts)
	if err != nil {
		panic(err)
	}

	// Send a ping to confirm a successful connection
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	Client.Client = client
	Client.Database = client.Database(DB_NAME)
}

type Document interface {
	Validate() error
	ValidatePreSave() error
	SetID(_id bson.ObjectID)
	GetID() bson.ObjectID
	SetCreatedAtAndUpdatedAt()
}

func Save[document Document](doc document, collection *mongo.Collection) error {

	doc.SetCreatedAtAndUpdatedAt()
	err := doc.ValidatePreSave()

	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := collection.InsertOne(ctx, doc)

	if err != nil {
		return err
	}

	_id, ok := res.InsertedID.(bson.ObjectID)

	if !ok {
		return fmt.Errorf("MongoError: ID is not an ObjectID, got type %T: %v", res.InsertedID, res.InsertedID)
	}

	doc.SetID(bson.ObjectID(_id))

	return doc.Validate()
}

func Update[document Document](doc document, collection *mongo.Collection, update bson.D) error {

	date_update := append(update, bson.D{{Key: "UpdatedAt", Value: time.Now()}}...)

	doc_fields := reflect.ValueOf(doc).Elem()

	for _, elem := range date_update {

		field := doc_fields.FieldByName(elem.Key)

		if field.IsValid() && field.CanSet() {
			value := reflect.ValueOf(elem.Value)

			if value.Type().AssignableTo(field.Type()) {
				field.Set(value)
			} else {
				return errors.New("INVALID UPDATE")
			}

		} else {
			return errors.New("INVALID UPDATE")
		}
	}

	err := doc.Validate()

	if err != nil {
		return err
	}

	full_update := bson.D{{Key: "$set", Value: date_update}}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = collection.UpdateOne(ctx, bson.D{{Key: "_id", Value: doc.GetID()}}, full_update)

	return err
}

func GetOne[document Document](filter bson.D, collection *mongo.Collection, empty_doc document) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := collection.FindOne(ctx, filter).Decode(empty_doc)

	if err != nil {
		return err
	}
	err = empty_doc.Validate()

	return err
}

func GetMany[document Document](filter bson.D, collection *mongo.Collection, empty_docs *[]document) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, filter)

	if err != nil {
		return err
	}

	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = cursor.All(ctx, empty_docs)

	if err != nil {
		return err
	}

	for _, doc := range *empty_docs {

		err = doc.Validate()
		if err != nil {
			return err
		}
	}

	return err
}
