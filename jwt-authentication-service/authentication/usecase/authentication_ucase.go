package usecase

import (
	"errors"
	domain "github.com/artmadar/jwt-auth-service/app-domain"
	"github.com/artmadar/jwt-auth-service/authentication/usecase/generator"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
	"log"
)

type authenticationUsecase struct {
	repo      domain.UserForAuthenticationRepository
	generator *generator.JwtGenerator
}

func NewAuthenticationUsecase(repo domain.UserForAuthenticationRepository, generator *generator.JwtGenerator) domain.AuthenticationUsecase {
	return &authenticationUsecase{
		repo,
		generator,
	}
}

func (a *authenticationUsecase) GenerateToken(ctx context.Context, username, password string) (*domain.Authentication, error) {
	if a.isCredentialsValid(ctx, username, password) {
		user, err := a.repo.GetUserByUsername(ctx, username)

		if err != nil {
			return nil, err
		}

		claims := map[string]interface{}{
			"username": user.Username,
			"ID":       user.ID,
			"email":    user.Email,
		}

		token, err := a.generator.GenerateJWT(claims)

		if err != nil {
			return nil, err
		}

		return &domain.Authentication{
			Token: token,
		}, nil

	}

	return nil, domain.ErrUnauthorized
}

func (a *authenticationUsecase) isCredentialsValid(ctx context.Context, username, password string) bool {
	user, err := a.repo.GetUserByUsername(ctx, username)

	if err != nil {
		log.Println("Getting user for validation error: ", err.Error())
		return false
	}

	if user == nil {
		log.Println("Somebody tried to authenticate with unexist username: ", username)
		return false
	}

	var ok, _ = a.passwordMatches(user.Password, password)

	return ok
}

// PasswordMatches uses Go's bcrypt package to compare a user supplied password
// with the hash we have stored for a given user in the database. If the password
// and hash match, we return true; otherwise, we return false.
func (a *authenticationUsecase) passwordMatches(password1, password2 string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(password1), []byte(password2))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			// invalid password
			return false, nil
		default:
			return false, err
		}
	}

	return true, nil
}
