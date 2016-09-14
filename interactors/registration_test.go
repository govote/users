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
	req.PhotoURL = "cadsa"

	interactor := new(RegistrationInteractor)
	interactor.UserRepository = new(FakeUserRepository)

	res := interactor.Register(req)

	assert.Equal(t, false, res.Success)
	assert.NotEmpty(t, res.Message)
}

func TestUserWithoutNameSholdReturnSuccessFalse(t *testing.T) {
	req := new(RegistrationRequest)
	req.Email = "test@test.com"
	req.FacebookID = "32132"
	req.PhotoURL = "cadsa"

	interactor := new(RegistrationInteractor)
	interactor.UserRepository = new(FakeUserRepository)

	res := interactor.Register(req)

	assert.Equal(t, false, res.Success)
	assert.NotEmpty(t, res.Message)
}

func TestUerWithoutFBIdSholdReturnSuccessFalse(t *testing.T) {
	req := new(RegistrationRequest)
	req.Email = "test@test.com"
	req.Name = "vitor"
	req.PhotoURL = "cadsa"

	interactor := new(RegistrationInteractor)
	interactor.UserRepository = new(FakeUserRepository)

	res := interactor.Register(req)

	assert.Equal(t, false, res.Success)
	assert.NotEmpty(t, res.Message)
}

func TestIfEmailExistsShouldReturnUserFromDatabase(t *testing.T) {
	email := "test@test.com"

	req := new(RegistrationRequest)
	req.Email = email
	req.Name = "vitor"
	req.FacebookID = "32132"
	req.PhotoURL = "cadsa"

	fakeRepository := new(FakeUserRepository)
	fakeRepository.On("FindByEmail", email).Return(domain.OptionalUser{Valid: true, User: domain.NewUser("registered", email)})

	interactor := new(RegistrationInteractor)
	interactor.UserRepository = fakeRepository

	res := interactor.Register(req)

	assert.Equal(t, true, res.Success)
	assert.NotNil(t, res.Data)
}

type FakeUserRepository struct{ mock.Mock }

func (fake *FakeUserRepository) Save(user domain.User) {}

func (fake *FakeUserRepository) FindByEmail(email string) domain.OptionalUser {
	args := fake.Called(email)
	return args.Get(0).(domain.OptionalUser)
}
