package interactors

import (
	"errors"

	"github.com/asaskevich/govalidator"
	"github.com/deputadosemfoco/users/domain"
	"github.com/deputadosemfoco/users/network"
)

type (
	AuthRequest struct {
		FacebookID          string `valid:"required"`
		FacebookAccessToken string `valid:"required"`
	}

	AuthResult struct {
		User          *domain.User
		Authenticated bool
	}

	AuthInteractorContract interface {
		Authenticate(req *AuthRequest) (*AuthResult, error)
	}

	AuthInteractor struct {
		UserRepository   domain.UserRepository
		FacebookGraphAPI network.FacebookGraphAPI
	}
)

func (interactor *AuthInteractor) Authenticate(req *AuthRequest) (*AuthResult, error) {
	val, err := govalidator.ValidateStruct(req)
	if !val {
		return nil, err
	}

	if !interactor.FacebookGraphAPI.ValidateAccessToken(req.FacebookAccessToken) {
		return &AuthResult{Authenticated: false}, nil
	}

	maybeUser := interactor.UserRepository.FindByFacebookID(req.FacebookID)

	if maybeUser.Valid {
		return &AuthResult{Authenticated: true, User: &maybeUser.User}, nil
	}

	return nil, errors.New("user not found")
}
