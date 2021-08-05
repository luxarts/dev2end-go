package repository

import (
	"clase2/defines"
	"clase2/domain"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

type UsersRepository interface {
	Create(user *domain.UserRepository) *domain.UserRepository
	GetByID(userID string) *domain.UserRepository
}

type usersRepository struct {
	mongoClient *mongo.Client
}

func NewUsersRepository() UsersRepository {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(defines.MongoDBURL))
	if err != nil {
		panic(err)
	}

	// Ping the primary
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	return &usersRepository{
		mongoClient: client,
	}
}

func (s *usersRepository) Create(user *domain.UserRepository) *domain.UserRepository{
	userBson := bson.D{
		{"id",user.ID},
		{"email",user.Email},
		{"name",user.Name},
		{"password",user.Password},
	}

	ctx, cancelCtx := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancelCtx()

	db := s.mongoClient.Database(defines.MongoDBTestDatabase)
	usersCollection := db.Collection(defines.MongoDBUsersCollection)
	resp, err := usersCollection.InsertOne(ctx, userBson)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", resp)

	userRepo := domain.UserRepository{}

	return &userRepo
}
func (s *usersRepository) GetByID(userID string) *domain.UserRepository {
	user := domain.UserRepository{}

	return &user
}