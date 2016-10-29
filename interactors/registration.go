package interactors

import (
	"errors"

	"github.com/asaskevich/govalidator"
	"github.com/deputadosemfoco/users/domain"
	"github.com/deputadosemfoco/users/network"
)

type (
	// RegistrationRequest represents the required information to register or login a User
	RegistrationRequest struct {
		Name                string `valid:"required"`
		Email               string `valid:"required,email"`
		PhotoURL            string `valid:"url,optional"`
		FacebookID          string `valid:"required"`
		FacebookAccessToken string `valid:"required"`
	}

	// RegistrationResult represents the result of registration interaction
	// Created will be true if a new user has been created in database
	RegistrationResult struct {
		Created bool
		User    domain.User
	}

	// RegistrationInteractorContract ...
	RegistrationInteractorContract interface {
		Register(req *RegistrationRequest) (*RegistrationResult, error)
	}

	// RegistrationInteractor is the interface for all registration operations
	RegistrationInteractor struct {
		UserRepository   domain.UserRepository
		FacebookGraphAPI network.FacebookGraphAPI
	}
)

// Register creates or simple login a User
func (interactor *RegistrationInteractor) Register(req *RegistrationRequest) (*RegistrationResult, error) {
	val, err := govalidator.ValidateStruct(req)
	if !val {
		return nil, err
	}

	if !interactor.FacebookGraphAPI.ValidateAccessToken(req.FacebookAccessToken) {
		return nil, errors.New("invalid facebook access token")
	}

	maybeUser := interactor.UserRepository.FindByEmail(req.Email)

	if maybeUser.Valid {
		return &RegistrationResult{Created: false, User: maybeUser.User}, nil
	}

	user := domain.NewUser(req.Name, req.Email, req.FacebookID)
	user.PhotoURL = req.PhotoURL

	return &RegistrationResult{Created: true, User: user}, nil
}
