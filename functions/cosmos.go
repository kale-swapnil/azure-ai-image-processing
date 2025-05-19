package cosmos

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Save(imageUrl string, tags []string) error {
	mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("COSMOSDB_URI")))
	defer client.Disconnect(context.TODO())

	doc := bson.M{"image_url": imageUrl, "tags": tags}
	_, err := client.Database("fashion").Collection("products").InsertOne(context.TODO(), doc)
	return err

}
