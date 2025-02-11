package services

import (
	"context"

	"github.com/Abdurahmanit/Game_log/backend/database"
	"github.com/Abdurahmanit/Game_log/backend/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var cartCollection *mongo.Collection = database.Client.Database("game-log").Collection("cart")

func AddToCart(cart model.Cart) error {
	_, err := cartCollection.InsertOne(context.TODO(), cart)
	return err
}

func GetCartByUserID(userID primitive.ObjectID) (model.Cart, error) {
	var cart model.Cart
	err := cartCollection.FindOne(context.TODO(), bson.M{"user_id": userID}).Decode(&cart)
	return cart, err
}
