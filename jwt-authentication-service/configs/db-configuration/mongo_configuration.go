package db_configuration

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"log"
	"time"
)

func GetMongoConnectedClient(mongoUrl, username, password string) (mongoClient *mongo.Client, cancel func()) {
	//connect to mongo
	mongoClient, err := connectToMongo(mongoUrl, username, password)

	if err != nil {
		log.Panic(err)
		return
	}
	log.Println("Service successfully connected to mongo-db: ", mongoUrl)

	//create context
	ctx, cl := context.WithTimeout(context.Background(), 15*time.Second)
	defer cl()

	cancel = func() {
		if err = mongoClient.Disconnect(ctx); err != nil {
			panic(err)
		}
		log.Println("Mongo client is going to be disconnected")
	}

	return
}

func connectToMongo(mongoUrl, username, password string) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(mongoUrl)
	clientOptions.SetAuth(options.Credential{
		Username: username,
		Password: password,
	})

	c, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Println("Error connection: ", err)
		return nil, err
	}

	return c, nil
}
