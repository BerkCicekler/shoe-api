package repository

import (
	"context"

	"github.com/BerkCicekler/shoe-api/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UsersRepo struct {
	MongoCollection *mongo.Collection	
}

func (r *UsersRepo) InsertUser(user *model.User) (interface{}, error) {
	result, err := r.MongoCollection.InsertOne(context.Background(), user)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *UsersRepo) FindUserByEmail(email string) (*model.User, error) {
	var user model.User
	filter := bson.D{{Key: "email", Value: email}}
	err := r.MongoCollection.FindOne(context.Background(), filter).Decode(&user)

	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UsersRepo) UpdateUserById(user *model.User) (int64, error) {
	result, err:= r.MongoCollection.UpdateOne(context.Background(), 
	bson.D{{Key: "_id", Value: user.ID}},
	bson.D{{Key: "$set", Value: user}})

	if err != nil {
		return 0, err 
	}
	return result.ModifiedCount, nil
}