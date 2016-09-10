package services

import (
	"fmt"

	"github.com/vitorsalgado/la-democracia/auth/domain/users"
	"github.com/vitorsalgado/la-democracia/lib/go/messages"
	"github.com/vitorsalgado/la-democracia/lib/go/validator"
)

type (
	// RegistrationService is the interface for all registration operations
	RegistrationService interface {
		Register(req *RegistrationRequest) messages.Response
	}

	// RegistrationServiceImpl is the default implementation for RegistrationService
	RegistrationServiceImpl struct {
	}
)

// NewRegistrationService returns a new instance of Service
func NewRegistrationService() *RegistrationServiceImpl {
	return &RegistrationServiceImpl{}
}

// Register creates or simple login a User
func (svc *RegistrationServiceImpl) Register(req *RegistrationRequest) messages.Response {
	var v = validator.Build()

	v.NotEmpty(req.Name, "name is required")
	v.NotEmpty(req.Email, "email is required")
	v.NotEmpty(req.FacebookID, "facebook ID is required")

	val, msgs := v.Validate()

	if !val {
		fmt.Println(fmt.Sprintf("invalid login attempt. err: %s", msgs))

		return messages.Response{Success: false, Message: "there are some invalid fields", ValidationMessages: msgs}
	}

	repository := users.NewUserRepository()
	optUser := repository.FindByEmail(req.Email)

	if optUser.Valid {
		user := optUser.User

		return messages.Response{Success: true, Data: user}
	}

	user := users.NewUser(req.Name, req.Email)

	user.FacebookID = req.FacebookID
	user.PhotoURL = req.PhotoURL

	repository.Save(user)

	return messages.Response{Success: true, Data: user}
}
