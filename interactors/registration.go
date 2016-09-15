package interactors

import (
	"github.com/deputadosemfoco/go-libs/messages"
	"github.com/deputadosemfoco/go-libs/validator"
	"github.com/deputadosemfoco/users/domain"
)

type (
	// RegistrationRequest represents the required information to register or login a User
	RegistrationRequest struct {
		messages.Request
		Name       string
		Email      string
		PhotoURL   string
		FacebookID string
	}

	RegistrationResult struct {
		messages.Response
		Created bool
	}

	// RegistrationInteractor is the interface for all registration operations
	RegistrationInteractor struct {
		UserRepository domain.UserRepository
	}
)

// Register creates or simple login a User
func (interactor *RegistrationInteractor) Register(req *RegistrationRequest) RegistrationResult {
	var v = validator.Build()

	v.NotEmpty(req.Name, "name is required")
	v.NotEmpty(req.Email, "email is required")
	v.NotEmpty(req.FacebookID, "facebook ID is required")

	val, msgs := v.Validate()

	if !val {
		res := messages.ResBuilder().AsErr("there are some invalid fields").WithValMsgs(msgs).Build()
		return RegistrationResult{Response: res}
	}

	optUser := interactor.UserRepository.FindByEmail(req.Email)

	if optUser.Valid {
		user := optUser.User

		return RegistrationResult{Response: messages.ResBuilder().AsSuccess(user).Build(), Created: false}
	}

	user := domain.NewUser(req.Name, req.Email)

	user.FacebookID = req.FacebookID
	user.PhotoURL = req.PhotoURL

	interactor.UserRepository.Save(user)

	return RegistrationResult{Response: messages.ResBuilder().AsSuccess(user).Build(), Created: true}
}
