package services

import (
	"Game-Log/backend/database"
	"Game-Log/backend/model"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var orderCollection *mongo.Collection = database.Client.Database("game-log").Collection("orders")

func CreateOrder(order model.Order) error {
	_, err := orderCollection.InsertOne(context.TODO(), order)
	return err
}

func GetOrdersByUserID(userID primitive.ObjectID) ([]model.Order, error) {
	var orders []model.Order
	cursor, err := orderCollection.Find(context.TODO(), bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var order model.Order
		cursor.Decode(&order)
		orders = append(orders, order)
	}

	return orders, nil
}
