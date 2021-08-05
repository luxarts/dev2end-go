package repository

import (
	"clase2/defines"
	"clase2/domain"
	"clase2/utils/jsend"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

type UsersRepository interface {
	Create(user *domain.UserRepository) (*domain.UserRepository, error)
	GetByID(userID string) (*domain.UserRepository, error)
	DeleteByID(userID string) error
}

type usersRepository struct {
	collection *mongo.Collection
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
		collection: client.Database(defines.MongoDBTestDatabase).Collection(defines.MongoDBUsersCollection),
	}
}

func (r *usersRepository) Create(user *domain.UserRepository) (*domain.UserRepository, error){
	userBson, err := bson.Marshal(user)
	if err != nil {
		return nil, jsend.NewError("marshal-error", err)
	}

	ctx, cancelCtx := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancelCtx()

	resp, err := r.collection.InsertOne(ctx, userBson)
	if err != nil {
		return nil, jsend.NewError("insertone-error", err)
	}

	// Guardo el ID con el que se creó
	user.ID = resp.InsertedID.(primitive.ObjectID)

	return user, nil
}
func (r *usersRepository) GetByID(userID string) (*domain.UserRepository, error) {
	ctx, cancelCtx := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancelCtx()

	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, jsend.NewError("objectidfromhex-error", err)
	}

	result := r.collection.FindOne(ctx, bson.M{"_id": objectID})
	if result.Err() == mongo.ErrNoDocuments {
		return nil, jsend.NewFailure("not-found")
	}

	var user domain.UserRepository
	if err := result.Decode(&user); err != nil {
		return nil, jsend.NewError("decode-error", err)
	}

	return &user, nil
}
func (r *usersRepository) DeleteByID(userID string) error {
	ctx, cancelCtx := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancelCtx()

	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return jsend.NewError("objectidfromhex-error", err)
	}

	result, err := r.collection.DeleteOne(ctx, bson.M{"_id": objectID})
	if result.DeletedCount == 0 {
		return jsend.NewFailure("not-found")
	}

	if err != nil {
		return jsend.NewError("deleteone-error", err)
	}

	return nil
}