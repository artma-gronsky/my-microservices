package main

import (
	"fmt"
	authenticationRepository "github.com/artmadar/jwt-auth-service/authentication/repository/mongo"
	authenticationUsecase "github.com/artmadar/jwt-auth-service/authentication/usecase"
	"github.com/artmadar/jwt-auth-service/authentication/usecase/generator"
	userHttpDelivery "github.com/artmadar/jwt-auth-service/user/delivery/http"
	"github.com/artmadar/jwt-auth-service/user/delivery/http/middleware"
	userRepo "github.com/artmadar/jwt-auth-service/user/repository/mongo"
	"github.com/artmadar/jwt-auth-service/user/repository/mongo/db-configuration"
	userUsecase "github.com/artmadar/jwt-auth-service/user/usecase"
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"time"
)

// todo: move to config
const (
	// todo: should be hidden somewhere
	jwtSecretKey     = "SecretYouShouldHide"
	webPort          = 1963
	timeoutInSeconds = 10
	mongoUrl         = "mongodb://localhost:27017"
	mongoUsername    = "mongoadmin"
	mongoPassword    = "secret"
)

type AppConfig struct {
}

func main() {
	app := AppConfig{}

	app.serve()
}

func (app *AppConfig) serve() {
	mux := chi.NewRouter()

	middL := middleware.InitMiddleware()
	mux.Use(middL.CORS)
	mux.Use(middL.Heartbeat)

	client, cancel := db_configuration.GetMongoConnectedClient(mongoUrl, mongoUsername, mongoPassword)
	defer cancel()

	ur := userRepo.NewMongoUserRepository(client)

	timeoutContext := time.Duration(timeoutInSeconds) * time.Second
	uc := userUsecase.NewUserUsecase(ur, timeoutContext)

	tr := authenticationRepository.NewMongoUserForAuthenticationRepository(client)
	ac := authenticationUsecase.NewAuthenticationUsecase(tr, generator.NewGenerator(jwtSecretKey))

	userHttpDelivery.NewUserHandler(mux, uc, ac)

	log.Println(fmt.Sprintf("JWT Authentication service is going to serve on port: %d\n", webPort))

	srv := http.Server{
		Addr:    fmt.Sprintf(":%d", webPort),
		Handler: mux,
	}
	srv.ListenAndServe()
}
