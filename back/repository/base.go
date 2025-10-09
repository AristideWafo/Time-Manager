package repository

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

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
