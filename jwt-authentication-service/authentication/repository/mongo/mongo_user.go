package mongo

import (
	domain "github.com/artmadar/jwt-auth-service/app-domain"
	"github.com/artmadar/jwt-auth-service/app-domain/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
)

type mongoUserRepository struct {
	client *mongo.Client
}

func NewMongoUserForAuthenticationRepository(mongo *mongo.Client) domain.UserForAuthenticationRepository {
	return &mongoUserRepository{client: mongo}
}

func (m *mongoUserRepository) GetUserByUsername(ctx context.Context, username string) (*entities.User, error) {
	collection := m.getUserCollection()

	result := collection.FindOne(ctx, bson.M{"username": username})

	var entry entities.User

	if err := result.Err(); err == mongo.ErrNoDocuments {
		return nil, nil
	}

	if err := result.Err(); err != nil {
		return nil, err
	}

	if err := result.Decode(&entry); err != nil {
		return nil, err
	}

	return &entry, nil
}

func (m *mongoUserRepository) getUserCollection() *mongo.Collection {
	collection := m.client.Database("jwtAuthenticationService").Collection("users")
	return collection
}
