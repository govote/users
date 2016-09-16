package interactors

import (
	"testing"

	"github.com/deputadosemfoco/users/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserWithoutEmailSholdReturnSuccessFalse(t *testing.T) {
	req := new(RegistrationRequest)
	req.Name = "vitor"
	req.FacebookID = "32132"
	req.FacebookAccessToken = "rewrw"
	req.PhotoURL = "http://photo.url.com.br"

	interactor := new(RegistrationInteractor)
	interactor.UserRepository = new(FakeUserRepository)
	interactor.FacebookGraphAPI = new(FakeFacebookGraphApi)

	res, err := interactor.Register(req)

	assert.NotNil(t, err)
	assert.Nil(t, res)
	assert.NotEmpty(t, err.Error())
}

func TestUserWithoutNameSholdReturnSuccessFalse(t *testing.T) {
	req := new(RegistrationRequest)
	req.Email = "test@test.com"
	req.FacebookID = "32132"
	req.FacebookAccessToken = "rewrw"
	req.PhotoURL = "http://photo.url.com.br"

	interactor := new(RegistrationInteractor)
	interactor.UserRepository = new(FakeUserRepository)
	interactor.FacebookGraphAPI = new(FakeFacebookGraphApi)

	res, err := interactor.Register(req)

	assert.Nil(t, res)
	assert.NotNil(t, err)
	assert.NotEmpty(t, err.Error())
}

func TestUerWithoutFBIdSholdReturnSuccessFalse(t *testing.T) {
	req := new(RegistrationRequest)
	req.Email = "test@test.com"
	req.Name = "vitor"
	req.FacebookAccessToken = "rewrw"
	req.PhotoURL = "http://photo.url.com.br"

	interactor := new(RegistrationInteractor)
	interactor.UserRepository = new(FakeUserRepository)
	interactor.FacebookGraphAPI = new(FakeFacebookGraphApi)

	res, err := interactor.Register(req)

	assert.NotNil(t, false, err)
	assert.Nil(t, res)
	assert.NotEmpty(t, err.Error())
}

func TestIfEmailExistsShouldReturnUserFromDatabase(t *testing.T) {
	email := "test@test.com"

	req := new(RegistrationRequest)
	req.Email = email
	req.Name = "vitor"
	req.FacebookID = "32132"
	req.FacebookAccessToken = "rewrw"
	req.PhotoURL = "http://photo.url.com.br"

	fakeRepository := new(FakeUserRepository)
	fakeRepository.On("FindByEmail", email).Return(domain.OptionalUser{Valid: true, User: domain.NewUser("registered", email, "fbid")})

	interactor := new(RegistrationInteractor)
	interactor.UserRepository = fakeRepository
	interactor.FacebookGraphAPI = new(FakeFacebookGraphApi)

	res, err := interactor.Register(req)

	assert.NotNil(t, res)
	assert.Nil(t, err)
	assert.NotNil(t, res.User)
}

type FakeUserRepository struct{ mock.Mock }
type FakeFacebookGraphApi struct{ mock.Mock }

func (fake *FakeUserRepository) Save(user domain.User) {}

func (fake *FakeUserRepository) FindByEmail(email string) domain.OptionalUser {
	args := fake.Called(email)
	return args.Get(0).(domain.OptionalUser)
}

func (fake *FakeFacebookGraphApi) ValidateAccessToken(accessToken string) bool {
	return true
}
