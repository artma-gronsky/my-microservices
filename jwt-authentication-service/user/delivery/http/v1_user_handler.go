package http

/*
It's api description for swagger

// CreateOrder godoc
// @Summary Create a new order
// @Description Create a new order with the input paylod
// @Tags orders
// @Accept  json
// @Produce  json
// @Param order body Order true "Create order"
// @Success 200 {object} Order
// @Router /orders [post]
*/

import (
	domain "github.com/artmadar/jwt-auth-service/app-domain"
	_ "github.com/artmadar/jwt-auth-service/user/delivery/http/docs"
	"github.com/artmadar/jwt-auth-service/user/delivery/http/helpers"
	"github.com/go-chi/chi"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

// UserHandler  represent the httpHandler for user
type userHandler struct {
	UserUsecase domain.UserUsecase
}

func NewUserHandler(mux *chi.Mux, us domain.UserUsecase) {
	handler := &userHandler{
		us,
	}

	//todo: remove hello method
	mux.Get("/", handler.hello)
	mux.Post("/api/v1/users", handler.store)

	mux.Mount("/swagger", httpSwagger.WrapHandler)
}

// Greet new user
// @Summary Greetings method
// @Description Greetings method (test method)
// @Tags greetings
// @Accept  json
// @Produce  json
// @Router / [get]
func (h *userHandler) hello(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("Hello world!"))
}

type UserPayload struct {
	ID          int64  `json:"id"`
	Username    string `json:"username" validate:"required"`
	Password    string `json:"password" validate:"required"`
	Email       string `json:"email" validate:"required"`
	Owner       string `json:"owner" validate:"required"`
	Description string `json:"description,omitempty"`
}

// Create new user
// @Summary Create a new user
// @Description Crete a new user in jwt-authentication-service
// @Tags user
// @Accept  json
// @Produce  json
// @Param order body UserPayload true "New user payload"
// @Success 200 {object} UserPayload
// @Router /api/v1/users [post]
func (h *userHandler) store(w http.ResponseWriter, r *http.Request) {
	var payload UserPayload
	err := helpers.ReadJSON(w, r, &payload)

	if err != nil {
		helpers.ErrorJSON(w, err)
	}
}
