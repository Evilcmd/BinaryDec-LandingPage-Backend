package mongodb

import (
	"context"
	"errors"

	"github.com/Evilcmd/Hackup-backend/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDbClient struct {
	MongoDbClient *mongo.Client
}

func NewMongoDbClient(uri string) (*MongoDbClient, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, errors.New("error creating a client: " + err.Error())
	}
	mongoClient := &MongoDbClient{client}
	return mongoClient, nil
}

func (mongoClient *MongoDbClient) FindUserWithEmail(email string) (models.User, error) {
	user := models.User{}

	filter := bson.D{{Key: "email", Value: email}}
	err := mongoClient.MongoDbClient.Database("binarydec").Collection("user").FindOne(context.Background(), filter).Decode(&user)

	return user, err
}

func (mongoClient *MongoDbClient) AddUser(user models.User) error {

	_, err := mongoClient.FindUserWithEmail(user.Email)
	if err == nil {
		return models.ErrorUserExists
	} else if err != mongo.ErrNoDocuments {
		return err
	}

	_, err = mongoClient.MongoDbClient.Database("binarydec").Collection("user").InsertOne(context.TODO(), user)

	return err
}
