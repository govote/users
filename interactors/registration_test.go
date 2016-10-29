package interactors

import (
	"testing"

	"github.com/deputadosemfoco/users/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserWithoutEmailSholdReturnError(t *testing.T) {
	req := new(RegistrationRequest)
	req.Name = "vitor"
	req.FacebookID = "32132"
	req.FacebookAccessToken = "valid"
	req.PhotoURL = "http://photo.url.com.br"

	graph := new(fakeRegFacebookGraphAPI)
	graph.On("ValidateAccessToken", req.FacebookAccessToken).Return(true)

	interactor := new(RegistrationInteractor)
	interactor.UserRepository = new(fakeRegUserRepository)
	interactor.FacebookGraphAPI = graph

	_, err := interactor.Register(req)

	assert.NotNil(t, err)
	assert.NotEmpty(t, err.Error())
}

func TestUserWithoutNameSholdReturnError(t *testing.T) {
	req := new(RegistrationRequest)
	req.Email = "test@test.com"
	req.FacebookID = "32132"
	req.FacebookAccessToken = "valid"
	req.PhotoURL = "http://photo.url.com.br"

	graph := new(fakeRegFacebookGraphAPI)
	graph.On("ValidateAccessToken", req.FacebookAccessToken).Return(true)

	interactor := new(RegistrationInteractor)
	interactor.UserRepository = new(fakeRegUserRepository)
	interactor.FacebookGraphAPI = graph

	_, err := interactor.Register(req)

	assert.NotNil(t, err)
	assert.NotEmpty(t, err.Error())
}

func TestUserWithoutFBIdSholdReturnError(t *testing.T) {
	req := new(RegistrationRequest)
	req.Email = "test@test.com"
	req.Name = "vitor"
	req.FacebookAccessToken = "valid"
	req.PhotoURL = "http://photo.url.com.br"

	graph := new(fakeRegFacebookGraphAPI)
	graph.On("ValidateAccessToken", req.FacebookAccessToken).Return(true)

	interactor := new(RegistrationInteractor)
	interactor.UserRepository = new(fakeRegUserRepository)
	interactor.FacebookGraphAPI = graph

	_, err := interactor.Register(req)

	assert.NotNil(t, err)
	assert.NotEmpty(t, err.Error())
}

func TestIfEmailExistsShouldReturnUserFromDatabase(t *testing.T) {
	email := "test@test.com"

	req := new(RegistrationRequest)
	req.Email = email
	req.Name = "vitor"
	req.FacebookID = "32132"
	req.FacebookAccessToken = "valid"
	req.PhotoURL = "http://photo.url.com.br"

	fakeRepository := new(fakeRegUserRepository)
	fakeRepository.On("FindByEmail", email).Return(domain.OptionalUser{Valid: true, User: domain.NewUser("registered", email, "fbid")})

	graph := new(fakeRegFacebookGraphAPI)
	graph.On("ValidateAccessToken", req.FacebookAccessToken).Return(true)

	interactor := new(RegistrationInteractor)
	interactor.UserRepository = fakeRepository
	interactor.FacebookGraphAPI = graph

	res, err := interactor.Register(req)

	assert.NotNil(t, res)
	assert.Nil(t, err)
	assert.NotNil(t, res.User)
}

func TestInvalidAccessTokenShouldReturnError(t *testing.T) {
	req := new(RegistrationRequest)
	req.Email = "test@test.com"
	req.Name = "vitor"
	req.FacebookAccessToken = "invalid"
	req.PhotoURL = "http://photo.url.com.br"

	graph := new(fakeRegFacebookGraphAPI)
	graph.On("ValidateAccessToken", req.FacebookAccessToken).Return(false)

	interactor := new(RegistrationInteractor)
	interactor.UserRepository = new(fakeRegUserRepository)
	interactor.FacebookGraphAPI = graph

	_, err := interactor.Register(req)

	assert.NotNil(t, err)
	assert.NotEmpty(t, err.Error())
}

func TestIfEmailDoesntExistsShouldReturnNewUser(t *testing.T) {
	req := new(RegistrationRequest)
	req.Email = "test@test.com"
	req.Name = "vitor"
	req.FacebookID = "32132"
	req.FacebookAccessToken = "valid"
	req.PhotoURL = "http://photo.url.com.br"

	fakeRepository := new(fakeRegUserRepository)
	fakeRepository.On("FindByEmail", req.Email).Return(domain.OptionalUser{Valid: false})

	graph := new(fakeRegFacebookGraphAPI)
	graph.On("ValidateAccessToken", req.FacebookAccessToken).Return(true)

	interactor := new(RegistrationInteractor)
	interactor.UserRepository = fakeRepository
	interactor.FacebookGraphAPI = graph

	res, err := interactor.Register(req)

	assert.NotNil(t, res)
	assert.Nil(t, err)
	assert.NotNil(t, res.User)
	assert.True(t, res.Created)
	assert.Equal(t, "test@test.com", res.User.Email)
	assert.Equal(t, "vitor", res.User.Name)
	assert.Equal(t, "32132", res.User.FacebookID)
	assert.Equal(t, "http://photo.url.com.br", res.User.PhotoURL)
}

type fakeRegUserRepository struct{ mock.Mock }
type fakeRegFacebookGraphAPI struct{ mock.Mock }

func (fake *fakeRegUserRepository) Save(user domain.User) {}

func (fake *fakeRegUserRepository) FindByEmail(email string) domain.OptionalUser {
	args := fake.Called(email)
	return args.Get(0).(domain.OptionalUser)
}

func (fake *fakeRegUserRepository) FindByFacebookID(facebookID string) domain.OptionalUser {
	args := fake.Called(facebookID)
	return args.Get(0).(domain.OptionalUser)
}

func (fake *fakeRegFacebookGraphAPI) ValidateAccessToken(accessToken string) bool {
	args := fake.Called(accessToken)

	return args.Get(0).(bool)
}
