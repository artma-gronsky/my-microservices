package main

import (
	"fmt"
	userHttpDelivary "github.com/artmadar/jwt-auth-service/user/delivery/http"
	"github.com/artmadar/jwt-auth-service/user/delivery/http/middleware"
	userRepo "github.com/artmadar/jwt-auth-service/user/repository/mongo"
	userUsecase "github.com/artmadar/jwt-auth-service/user/usecase"
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"time"
)

// todo: move to config
const (
	webPort          = 1963
	timeoutInSeconds = 2
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

	ur := userRepo.NewMongoUserRepository()

	timeoutContext := time.Duration(timeoutInSeconds) * time.Second
	uc := userUsecase.NewUserUsecase(ur, timeoutContext)

	userHttpDelivary.NewUserHandler(mux, uc)

	log.Println(fmt.Sprintf("JWT Authentication service is going to serve on port: %d\n", webPort))

	srv := http.Server{
		Addr:    fmt.Sprintf(":%d", webPort),
		Handler: mux,
	}
	srv.ListenAndServe()
}
