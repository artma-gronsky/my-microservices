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

func (m *mongoUserRepository) GetOneByOneOfFieldsValues(ctx context.Context, fieldsValues map[string]string) (*domain.User, error) {
	collection := m.getUserCollection()

	//result := collection.FindOne(ctx, bson.M{}, options.FindOne().SetProjection(bson.M{"username": username}))
	//result := collection.FindOne(ctx, bson.M{"username": username, "email": "Alex.artmadar@gmail.com"})
	//bson.D{{"a": 1}, {"b": 2}}

	pipeline := make([]interface{}, 0)
	for k, v := range fieldsValues {
		regex := bson.M{"$regex": primitive.Regex{Pattern: v, Options: "i"}}
		pipeline = append(pipeline, bson.D{{k, regex}})
	}
	result := collection.FindOne(ctx, bson.D{
		{"$or", pipeline},
	})

	var user domain.User

	switch result.Err() {
	case mongo.ErrNoDocuments:
		return nil, nil
	case nil:
		break
	default:
		return nil, result.Err()
	}

	if err := result.Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}
func (m *mongoUserRepository) GetById(ctx context.Context, idHex string) (*domain.User, error) {
	collection := m.getUserCollection()

	objectId, err := primitive.ObjectIDFromHex(idHex)

	if err != nil {
		return nil, err
	}

	result := collection.FindOne(ctx, bson.M{"_id": objectId})

	var entry domain.User

	if err := result.Err(); err != nil {
		return nil, err
	}

	err = result.Decode(&entry)

	if err != nil {
		return nil, err
	}

	return &entry, nil
}

func (m *mongoUserRepository) getUserCollection() *mongo.Collection {
	collection := m.client.Database("jwtAuthenticationService").Collection("users")
	return collection
}
