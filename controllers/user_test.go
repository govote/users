package controllers

import (
	"errors"
	"net/http"
	"testing"

	"github.com/deputadosemfoco/go-libs/test"
	"github.com/deputadosemfoco/users/interactors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestWhenCreatingNewUserStatusCodeMustBeCreated201(t *testing.T) {
	// setup
	param := new(interactors.RegistrationRequest)
	param.Email = "test@test.com"
	param.Name = "vitor"
	param.FacebookID = "32132"
	param.PhotoURL = "cadsa"

	userCtrl := new(UserCtrl)
	fakeInteractor := new(FakeRegistrationInteractor)

	fakeInteractor.On("Register", param).Return(&interactors.RegistrationResult{Created: true}, nil)
	userCtrl.Interactor = fakeInteractor

	c, res, err := test.CreateJSONContext(param)

	test.PanicErr(err)

	// execution
	err = userCtrl.Post(c)

	// assertions
	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, res.Code)
}

func TestUserWithInvalidInformationShouldReturnBadRequest400(t *testing.T) {
	// setup
	param := new(interactors.RegistrationRequest)
	param.Name = "vitor"
	param.FacebookID = "32132"
	param.PhotoURL = "cadsa"

	userCtrl := new(UserCtrl)
	fakeInteractor := new(FakeRegistrationInteractor)

	fakeInteractor.On("Register", param).Return(&interactors.RegistrationResult{}, errors.New("error"))
	userCtrl.Interactor = fakeInteractor

	c, res, err := test.CreateJSONContext(param)

	test.PanicErr(err)

	// execution
	err = userCtrl.Post(c)

	// assertions
	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, res.Code)
}

func TestAlreadyRegisteredUserShouldReturnOk200(t *testing.T) {
	// setup
	param := new(interactors.RegistrationRequest)
	param.Email = "test@test.com"
	param.Name = "vitor"
	param.FacebookID = "32132"
	param.PhotoURL = "cadsa"

	userCtrl := new(UserCtrl)
	fakeInteractor := new(FakeRegistrationInteractor)

	fakeInteractor.On("Register", param).Return(&interactors.RegistrationResult{Created: false}, nil)
	userCtrl.Interactor = fakeInteractor

	c, res, err := test.CreateJSONContext(param)

	test.PanicErr(err)

	// execution
	err = userCtrl.Post(c)

	// assertions
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.Code)
}

func TestEmptyJSONShouldReturnError(t *testing.T) {
	userCtrl := new(UserCtrl)
	c, _, err := test.CreateJSONContext("")

	test.PanicErr(err)

	err = userCtrl.Post(c)

	assert.NotNil(t, err)
}

type FakeRegistrationInteractor struct{ mock.Mock }

func (fake *FakeRegistrationInteractor) Register(request *interactors.RegistrationRequest) (*interactors.RegistrationResult, error) {
	args := fake.Called(request)

	return args.Get(0).(*interactors.RegistrationResult), args.Error(1)
}
