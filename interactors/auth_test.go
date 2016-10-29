package interactors

import (
	"testing"

	"github.com/deputadosemfoco/users/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAuthWithoutFBIdSholdReturnError(t *testing.T) {
	req := new(AuthRequest)
	req.FacebookAccessToken = "valid"
	req.FacebookID = ""

	graph := new(fakeFacebookGraphAPI)
	graph.On("ValidateAccessToken", req.FacebookAccessToken).Return(true)

	repository := new(fakeUserRepository)
	repository.On("FindByFacebookID", req.FacebookID).Return(&domain.OptionalUser{Valid: false})

	interactor := new(AuthInteractor)
	interactor.UserRepository = repository
	interactor.FacebookGraphAPI = graph

	_, err := interactor.Authenticate(req)

	assert.NotNil(t, err)
	assert.NotEmpty(t, err.Error())
}

type fakeUserRepository struct{ mock.Mock }
type fakeFacebookGraphAPI struct{ mock.Mock }

func (fake *fakeUserRepository) Save(user domain.User) {}

func (fake *fakeUserRepository) FindByEmail(email string) domain.OptionalUser {
	args := fake.Called(email)
	return args.Get(0).(domain.OptionalUser)
}

func (fake *fakeUserRepository) FindByFacebookID(facebookID string) domain.OptionalUser {
	args := fake.Called(facebookID)
	return args.Get(0).(domain.OptionalUser)
}

func (fake *fakeFacebookGraphAPI) ValidateAccessToken(accessToken string) bool {
	args := fake.Called(accessToken)

	return args.Get(0).(bool)
}
