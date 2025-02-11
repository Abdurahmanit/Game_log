package services

import (
	"Game-Log/backend/database"
	"Game-Log/backend/model"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.Client.Database("game-log").Collection("users")

func RegisterUser(user model.User) error {
	_, err := userCollection.InsertOne(context.TODO(), user)
	return err
}

func LoginUser(email, password string) (model.User, error) {
	var user model.User
	err := userCollection.FindOne(context.TODO(), bson.M{"email": email, "password": password}).Decode(&user)
	return user, err
}
