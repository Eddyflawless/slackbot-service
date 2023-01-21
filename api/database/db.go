package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var dbClient *mongo.Client

func CreateDBConnection() {

	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGODB_URL")))

	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := CreateTTLContext()

	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to database")

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	dbClient = client
}

func OpenCollection(collectionName string) *mongo.Collection {

	database := os.Getenv("MONGO_DB")

	return dbClient.Database(database).Collection(collectionName)

}

func ConvertObjectIDToHex(id string) (objectId primitive.ObjectID, err error) {

	objectId, err = primitive.ObjectIDFromHex(id)

	if err != nil {
		log.Println("Invalid id")
		return objectId, err
	}

	return objectId, nil

}

func CreateTTLContext() (ctx context.Context, cancel context.CancelFunc) {

	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	return ctx, cancel

}

func FindOne(collectionName string, entityModel interface{}, filter interface{}, opts *options.FindOneOptions) (err error) {

	ctx, cancel := CreateTTLContext()

	err = OpenCollection(collectionName).FindOne(ctx, filter, opts).Decode(entityModel)

	defer cancel()

	return err
}

func FindMany(collectionName string, filter interface{}, opts *options.FindOptions) (cursor *mongo.Cursor, err error) {

	ctx, cancel := CreateTTLContext()

	cursor, err = OpenCollection(collectionName).Find(ctx, filter, opts)

	defer cancel()

	return cursor, err

}

func UpdateOne(collectionName string, filter interface{}, update interface{}, opts *options.UpdateOptions) (result *mongo.UpdateResult, err error) {

	ctx, cancel := CreateTTLContext()

	result, err = OpenCollection(collectionName).UpdateOne(ctx, filter, update)

	defer cancel()

	return result, err
}
