package services

import (
	"context"

	"github.com/Abdurahmanit/Game_log/backend/database"
	"github.com/Abdurahmanit/Game_log/backend/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var gameCollection *mongo.Collection = database.Client.Database("game-log").Collection("games")

func GetAllGames() ([]model.Game, error) {
	var games []model.Game
	cursor, err := gameCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var game model.Game
		cursor.Decode(&game)
		games = append(games, game)
	}

	return games, nil
}
