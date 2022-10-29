package usecase

import (
	"context"
	domain "github.com/artmadar/jwt-auth-service/app-domain"
	"github.com/artmadar/jwt-auth-service/app-domain/entities"
	"golang.org/x/crypto/bcrypt"
	"log"
	"strings"
	"time"
)

type userUsecase struct {
	userRepo       domain.UserRepository
	contextTimeout time.Duration
}

// NewUserUsecase will create new an userUsecase object representation of domain.UserUsecase interface
func NewUserUsecase(a domain.UserRepository, timeout time.Duration) domain.UserUsecase {
	return &userUsecase{
		userRepo:       a,
		contextTimeout: timeout,
	}
}

func (u *userUsecase) Store(ctx context.Context, createUserRequest *entities.User) (err error) {

	Validate(createUserRequest)

	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	existUser, err := u.findEntryWithEmailOrUsername(ctx, createUserRequest.Username, createUserRequest.Email)

	if err != nil {
		log.Println(err.Error())
		return err
	}

	if existUser != nil {
		if strings.ToLower(createUserRequest.Email) == strings.ToLower(existUser.Email) {
			return domain.ErrUserWithEmailAlreadyExist
		}

		if strings.ToLower(createUserRequest.Username) == strings.ToLower(existUser.Username) {
			return domain.ErrUserWithUsernameAlreadyExist
		}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(createUserRequest.Password), 12)

	err = u.userRepo.Store(ctx, &entities.User{
		Username:    createUserRequest.Username,
		Password:    string(hashedPassword),
		Email:       createUserRequest.Email,
		Description: createUserRequest.Description,
		CreatedAt:   time.Now().UTC(),
	})

	return
}

func (u *userUsecase) findEntryWithEmailOrUsername(ctx context.Context, username, email string) (*entities.User, error) {
	fieldsValues := map[string]string{
		"username": username,
		"email":    email,
	}
	return u.userRepo.GetOneByOneOfFieldsValues(ctx, fieldsValues)
}
