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
// @Param        username    query     string  true "user's username"
// @Param 		 email    query 	   string  false "user's email"
// @Success 200 {object} Order
// @Router /orders [post]
*/

import (
	domain "github.com/artmadar/jwt-auth-service/app-domain"
	"github.com/artmadar/jwt-auth-service/app-domain/entities"
	_ "github.com/artmadar/jwt-auth-service/user/delivery/http/docs"
	"github.com/artmadar/jwt-auth-service/user/delivery/http/helpers"
	"github.com/go-chi/chi"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

// UserHandler  represent the httpHandler for user
type userHandler struct {
	UserUsecase           domain.UserUsecase
	AuthenticationUsecase domain.AuthenticationUsecase
}

func NewUserHandler(mux *chi.Mux, us domain.UserUsecase, au domain.AuthenticationUsecase) {
	handler := &userHandler{
		us,
		au,
	}

	//todo: remove hello method
	mux.Get("/", handler.hello)
	mux.Post("/api/v1/users", handler.store)
	mux.Post("/api/v1/users/token", handler.token)

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
	Username    string `json:"username" validate:"required"`
	Password    string `json:"password" validate:"required"`
	Email       string `json:"email" validate:"required"`
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
		return
	}

	ctx := r.Context()
	err = h.UserUsecase.Store(ctx, &entities.User{
		Username:    payload.Username,
		Password:    payload.Password,
		Email:       payload.Email,
		Description: payload.Description,
	})

	if err != nil {
		helpers.ErrorJSON(w, err)
		return
	}

	helpers.WriteJson(w, http.StatusAccepted, payload)
}

type Credentials struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type TokenPayload struct {
	Token string `json:"token" validate:"required"`
}

// Get user's token
// @Summary Get token for user
// @Description Generate a new access token for user
// @Tags user
// @Accept  json
// @Produce  json
// @Param credentials body Credentials true "User credentials"
// @Success 200 {object} TokenPayload
// @Router /api/v1/users/token [post]
func (h *userHandler) token(w http.ResponseWriter, r *http.Request) {
	var payload Credentials
	err := helpers.ReadJSON(w, r, &payload)

	if err != nil {
		helpers.ErrorJSON(w, err)
		return
	}

	ctx := r.Context()

	responsePayload, err := h.AuthenticationUsecase.GenerateToken(ctx, payload.Username, payload.Password)

	if err != nil {
		helpers.ErrorJSON(w, err)
		return
	}

	helpers.WriteJson(w, http.StatusOK, responsePayload)

}
