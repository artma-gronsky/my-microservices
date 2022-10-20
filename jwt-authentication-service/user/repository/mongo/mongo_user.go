package mongo

import (
	"context"
	domain "github.com/artmadar/jwt-auth-service/app-domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type mongoUserRepository struct {
	client *mongo.Client
}

func NewMongoUserRepository(mongo *mongo.Client) domain.UserRepository {
	return &mongoUserRepository{client: mongo}
}

func (m *mongoUserRepository) Store(ctx context.Context, user *domain.User) error {
	_, err := m.getUserCollection().InsertOne(ctx, user)

	if err != nil {
		log.Println("Error inserting on the users:", err.Error())
		return err
	}

	return nil
}

func (m *mongoUserRepository) GetById(ctx context.Context, idHex string) (*domain.User, error) {

	objectId, err := primitive.ObjectIDFromHex(idHex)

	if err != nil {
		return nil, err
	}

	result := m.getUserCollection().FindOne(ctx, bson.M{"_id": objectId})

	var user domain.User

	if err := result.Err(); err != nil {
		return nil, err
	}

	err = result.Decode(user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (m *mongoUserRepository) getUserCollection() *mongo.Collection {
	collection := m.client.Database("jwtAuthenticationService").Collection("users")
	return collection
}
