package controller

import (
	"context"
	"fmt"
	"log"

	"github.com/mihir010/DbApi/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://kumarmihir02:DlRzvLrUtXEoUVcJ@gofirst.sj5svf3.mongodb.net/?retryWrites=true&w=majority"
const dbName = "netflix"
const colName = "watchlist"

var collection *mongo.Collection

func init() {
	clientOption := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database(dbName).Collection(colName)
}

func insertData(show models.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), show)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Inserted 1 show with show ID: %d", inserted.InsertedID)
}

func updateData(showId string) {
	id, _ := primitive.ObjectIDFromHex(showId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("successfully updated: ", result.ModifiedCount)
}

func deleteData(showId string) {
	id, _ := primitive.ObjectIDFromHex(showId)
	filter := bson.M{"_id": id}
	result, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("deleted successfully ", result.DeletedCount)
}
