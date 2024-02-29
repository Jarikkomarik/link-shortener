package repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"com.jarikkomarik.linkshortener/myError"
)

const collectionName = "link-collection"

type LinkRepository struct {
	linkCollection *mongo.Collection
}

func NewLinkRepository(linkDatabase *mongo.Database) *LinkRepository {
	return &LinkRepository{linkCollection: linkDatabase.Collection(collectionName)}
}

func (linkRepository *LinkRepository) InsertRecord(url string) string {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	generatedId, err := linkRepository.linkCollection.InsertOne(ctx, bson.M{"url": url})
	if err != nil {
		panic(err)
	}
	return generatedId.InsertedID.(primitive.ObjectID).Hex()
}

func (linkRepository *LinkRepository) GetUrlId(url string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	// FindOne returns a single document matching the filter
	result := linkRepository.linkCollection.FindOne(ctx, bson.D{{"url", url}})

	// Decode the result into a string
	var id map[string]string
	if err := result.Decode(&id); err != nil {
		return "", err
	}

	return id["_id"], nil
}

func (linkRepository *LinkRepository) GetRecord(id string) string {
	var result map[string]string
	val, objectIdErr := primitive.ObjectIDFromHex(id)

	if objectIdErr != nil {
		panic(myError.InvalidRecordId{})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	err := linkRepository.linkCollection.FindOne(ctx, bson.M{"_id": val}).Decode(&result)
	if err != nil {
		panic(err)
	}
	return result["url"]
}
